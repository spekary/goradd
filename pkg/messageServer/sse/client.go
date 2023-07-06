// Portions copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sse

import (
	"encoding/json"
	log2 "github.com/goradd/goradd/pkg/log"
	"net/http"
)

// client keeps track of client specific information.
type client struct {
	hub *ServerSentEventHub

	// The server sent event connection
	w http.ResponseWriter

	// Buffered channel of outbound messages.
	send chan clientMessage

	// channels are the channel ids that this client is subscribed to
	channels map[string]struct{}

	clientID string // authenticator
}

func newClient(hub *ServerSentEventHub, w http.ResponseWriter, r *http.Request, clientID string) {

	log2.FrameworkDebugf("New client %s", clientID)

	c := client{hub: hub,
		w:        w,
		clientID: clientID,
		send:     make(chan clientMessage, 16),
		channels: make(map[string]struct{}),
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	done := r.Context().Done()

	go func() {
		<-done
		hub.unregister <- clientID
	}()
	hub.register <- &c

	go c.writePump()
}

// writePump pumps messages from the hub to the outbound connections.
//
// A goroutine running writePump is started for each client. The
// application ensures that there is at most one writer to a client by
// executing all writes from this goroutine.
func (c *client) writePump() {
	for {
		message, ok := <-c.send
		if !ok {
			// The hub closed the client.
			return // end the loop
		}
		// messages are json objects. We gather them up here into an array.
		messages := []clientMessage{message}

		// Add queued messages to the current message.
		n := len(c.send)
		for i := 0; i < n; i++ {
			messages = append(messages, <-c.send)
		}

		if buf, err2 := json.Marshal(messages); err2 != nil {
			panic(err2) // malformed message. Shouldn't happen.
		} else {
			if _, err := c.w.Write(buf); err != nil {
				return
			}
			log2.FrameworkDebugf("Writepump %s", string(buf))
		}

		if f, ok := c.w.(http.Flusher); ok {
			f.Flush()
		}
	}
}
