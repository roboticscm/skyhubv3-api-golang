package main

import (
	"backend/system/config"
	"backend/system/db"
	"backend/system/server"
	"backend/system/slog"
	"fmt"
	"time"

	"github.com/lib/pq"
)

var app server.Server

func init() {
	// load config
	conf, err := config.LoadCommonConfig()
	if err != nil {
		slog.Fatal("Common Config Error", err)
	}
	// init dbBegooDB
	beeGoDB := db.BeegoDB{}
	beeGoDB.Init(conf)
	// sync db
	// beeGoDB.Sync()

	// Start HTTP Sever
	app = server.Server{
		Port: conf.ServerPort,
	}
}

func notifyListener(b *db.Broker) {
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(db.DBConnectionStr, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen("event_channel")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Start monitoring PostgreSQL...")
	for {
		db.WaitForNotification(listener, b)
	}
}

func main() {
	app.Init()
	b := &db.Broker{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}
	b.Start()
	// notify listener postgres
	go notifyListener(b)
	app.RegisterRoute(b)
	app.Start()
}
