// Portions copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sse

import (
	"bytes"
	"encoding/json"
	log2 "github.com/goradd/goradd/pkg/log"
	"net/http"
	"time"
)

// client keeps track of client specific information.
type client struct {
	hub *ServerSentEventHub

	// The server sent event connection
	w http.ResponseWriter
	r *http.Request

	// Buffered channel of outbound messages.
	send chan clientMessage

	// channels are the channel ids that this client is subscribed to
	channels map[string]struct{}

	clientID string // authenticator
}

func newClient(hub *ServerSentEventHub, w http.ResponseWriter, r *http.Request, clientID string) {

	log2.FrameworkDebugf("New messenger client %s", clientID)

	c := client{hub: hub,
		w:        w,
		clientID: clientID,
		send:     make(chan clientMessage, 16),
		channels: make(map[string]struct{}),
		r:        r,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	rc := http.NewResponseController(w)
	if err := rc.SetWriteDeadline(time.Now().Add(WriteTimeout)); err != nil {
		log2.FrameworkDebug("Response Controller error: ", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	msg := ":ka\n\n" // Safari and perhaps other browsers will not open the connection until the first message is
	// attempted.
	if _, err := c.w.Write([]byte(msg)); err != nil {
		return
	}
	if err := c.flush(); err != nil {
		return
	}

	hub.register <- &c

	defer hub.unregisterClient(c.clientID)

	c.writePump()

}

// writePump pumps messages from the hub to the outbound connections.
//
// writePump does not return until the connection is closed, and so holds on to the
// connection until that happens.
func (c *client) writePump() {
	ticker := time.NewTicker(KeepAliveInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.r.Context().Done():
			return

		case message, ok := <-c.send:
			if !ok || message.Message == "" {
				// The hub closed the client.
				log2.Infof("Connection closed")
				return
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
				var buf2 bytes.Buffer
				buf2.WriteString("data: ")
				buf2.Write(buf)
				buf2.WriteString("\n\n")

				if _, err := c.w.Write(buf2.Bytes()); err != nil {
					log2.FrameworkDebug("Writepump error: ", err.Error())
					return // closes the connection
				}
				log2.FrameworkDebugf("Writepump - clientID: %s, msg: %s", c.clientID, string(buf))
			}
		case <-ticker.C:
			msg := ":ka\n\n" // a keepalive comment
			if _, err := c.w.Write([]byte(msg)); err != nil {
				log2.FrameworkDebug("Ticker error: ", err.Error())
				return // closes the connection
			}
			log2.FrameworkDebugf("Keep alive: %s", c.clientID)
		}

		if err := c.flush(); err != nil {
			return // closes the connection
		}
	}
}

func (c *client) flush() (err error) {
	rc := http.NewResponseController(c.w)
	err = rc.Flush()
	if err != nil {
		log2.FrameworkDebug("Flush error: ", err.Error())
	}
	return
}

func (c *client) close() {
	c.hub.unregisterClient(c.clientID)
}
