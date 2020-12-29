package db

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

func WaitForNotification(l *pq.Listener) {
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
			fmt.Println(string(prettyJSON.Bytes()))
			return
			// case <-time.After(90 * time.Second):
			// 	fmt.Println("Received no events for 90 seconds, checking connection")
			// 	go func() {
			// 		l.Ping()
			// 	}()
			// 	return
		}
	}
}
