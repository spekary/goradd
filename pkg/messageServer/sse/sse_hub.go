// Portions copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sse

import (
	"github.com/goradd/goradd/pkg/log"
	"time"
)

type clientMessage struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

type subscription struct {
	clientID string
	channel  string
}

// KeepAliveInterval is the interval that the server will ping the client to keep it from closing the connection.
var KeepAliveInterval = time.Duration(10 * time.Minute)

// WriteTimeout is the total time that the connection can be kept open. Browsers will need to refresh their pages
// to re-establish the connection after this time. This should be kept as small as possible to prevent running out
// of resources.
var WriteTimeout = time.Duration(24 * time.Hour)

type ServerSentEventHub struct {
	// Channels that clients have subscribed to. Each channel points to a map of client IDs
	channels map[string]map[string]struct{}

	// Registered clients, keyed by client ID
	clients map[string]*client

	// Register requests from the clients.
	register chan *client

	// Unregister requests from clients.
	unregister chan string

	send chan clientMessage

	subscribe chan subscription
}

func NewServerSentEventHub() *ServerSentEventHub {
	return &ServerSentEventHub{
		register:   make(chan *client),
		unregister: make(chan string),
		send:       make(chan clientMessage),
		clients:    make(map[string]*client),
		channels:   make(map[string]map[string]struct{}),
		subscribe:  make(chan subscription),
	}
}

func (h *ServerSentEventHub) run() {
	for {
		select {
		case client := <-h.register:
			log.Infof("SSE client registering - client ID: %s", client.clientID)

			if _, ok := h.clients[client.clientID]; ok {
				// The same client is registering again. Unregister first.
				h.unregisterClient(client.clientID)
			}
			h.clients[client.clientID] = client

		case clientID := <-h.unregister:
			log.Infof("SSE client unregistering - clientID: %s", clientID)
			h.unregisterClient(clientID)

		case msg := <-h.send:
			if clientIDs, ok := h.channels[msg.Channel]; ok {
				log.FrameworkDebugf("Sending to channel %s - %v", msg.Channel, msg.Message)

				for clientID := range clientIDs {
					if client, ok2 := h.clients[clientID]; ok2 {
						log.FrameworkDebugf("Sending to client: %s", client.clientID)
						client.send <- msg
					}
				}
			} else {
				// Sending to a channel that has not been allocated, so ignore it
				log.Errorf("Could not find channel %s", msg.Channel)
			}

		case sub := <-h.subscribe:
			h.subscribeChannel(sub.clientID, sub.channel)
		}
	}
}

func (h *ServerSentEventHub) unregisterClient(clientID string) {
	log.FrameworkDebugf("UnregisterClient %s", clientID)

	var client, _ = h.clients[clientID]

	if client == nil {
		return
	}

	for channel := range client.channels {
		h.unsubscribeChannel(clientID, channel)
	}
	delete(h.clients, clientID)
}

func (h *ServerSentEventHub) subscribeChannel(clientID string, channel string) {
	log.FrameworkDebugf("Subscribe client: %s channel: %s", clientID, channel)

	var client, _ = h.clients[clientID]

	if client == nil {
		return
	}

	client.channels[channel] = struct{}{}
	if clientIDs, ok := h.channels[channel]; !ok {
		clientIDs = make(map[string]struct{})
		clientIDs[clientID] = struct{}{}
		h.channels[channel] = clientIDs
	} else {
		clientIDs[clientID] = struct{}{}
	}
}

func (h *ServerSentEventHub) unsubscribeChannel(clientID string, channel string) {
	log.FrameworkDebugf("Unsubscribe client: %s channel: %s", clientID, channel)

	if clientIDs, ok := h.channels[channel]; ok {
		delete(clientIDs, clientID)
		if len(clientIDs) == 0 {
			delete(h.channels, channel)
		} else {
			h.channels[channel] = clientIDs
		}
	}
	if client, ok := h.clients[clientID]; ok {
		delete(client.channels, channel)
	}
}
