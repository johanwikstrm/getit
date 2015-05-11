package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
)

func DataFeed(sess *session) func(*websocket.Conn) {
	fmt.Println("Started another datafeed for session ", sess.shorturl)
	return func(ws *websocket.Conn) {
		for {
			sess.newMsg.L.Lock()
			// send new message
			//fmt.Println(ws, "Sending")
			msg := sess.countvotes()
			if err := websocket.JSON.Send(ws, &msg); err != nil {
				log.Println(err)
				sess.newMsg.L.Unlock()
				break
			}
			//fmt.Println(ws, "Waiting")
			sess.newMsg.Wait() // wait for new message
			//fmt.Println(ws, "Done waiting")
			sess.newMsg.L.Unlock()
		}
	}
}
