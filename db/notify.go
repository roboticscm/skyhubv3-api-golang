package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type Broker struct {
	Clients        map[chan string]bool
	NewClients     chan chan string
	DefunctClients chan chan string
	Messages       chan string
}

func (b *Broker) Start() {
	go func() {
		for {
			select {
			case s := <-b.NewClients:
				b.Clients[s] = true
				log.Println("Added new client")

			case s := <-b.DefunctClients:
				delete(b.Clients, s)
				close(s)
				log.Println("Removed client")

			case msg := <-b.Messages:
				for s := range b.Clients {
					s <- msg
				}
				log.Printf("Broadcast message to %d clients", len(b.Clients))
			}
		}
	}()
}

func (b *Broker) NotifyHandler(c echo.Context) error {
	w := c.Response().Writer
	f, ok := w.(http.Flusher)
	if !ok {
		return errors.New("Streaming unsupported!")
	}
	messageChan := make(chan string)
	b.NewClients <- messageChan

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		b.DefunctClients <- messageChan
		log.Println("HTTP connection just closed.")
	}()

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Transfer-Encoding", "chunked")
	c.Response().Header().Set("Access-Control-Allow-Credentials", "false")
	c.Response().Header().Set("Access-Control-Allow-Headers", "x-requested-with, authorization, Content-Type, Authorization, X-XSRF-TOKEN")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	c.Response().WriteHeader(http.StatusOK)

	for {
		msg, open := <-messageChan
		if !open {
			break
		}
		fmt.Fprintf(w, "data:%s\n\n", msg)
		f.Flush()
	}

	return nil
}

func WaitForNotification(l *pq.Listener, b *Broker) {
	for {
		select {
		case n := <-l.Notify:
			fmt.Println("Received data from channel [", n.Channel, "] :")
			// Prepare notification payload for pretty print
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
			if err != nil {
				fmt.Println("Error processing JSON: ", err)
				return
			}

			b.Messages <- strings.ReplaceAll(string(prettyJSON.Bytes()), "\n", "")
			return
		case <-time.After(90 * time.Second):
			fmt.Println("Received no events for 90 seconds, checking connection")
			go func() {
				l.Ping()
			}()
			return
		}
	}
}
