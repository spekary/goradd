package sse

import (
	"fmt"
	"github.com/goradd/goradd/pkg/config"
	http2 "github.com/goradd/goradd/pkg/http"
	_ "github.com/goradd/goradd/pkg/messageServer/sse/assets"
	"github.com/goradd/html5tag"
	"net/http"
	"path"
)

type SseMessenger struct {
	hub *ServerSentEventHub
}

func (m *SseMessenger) Start() *ServerSentEventHub {
	m.hub = NewServerSentEventHub()
	go m.hub.run()
	return m.hub
}

func (m *SseMessenger) JavascriptInit() string {
	return fmt.Sprintf("goradd.initMessagingClient(%q);\n", http2.MakeLocalPath(config.MessengerPrefix))
}

func (m *SseMessenger) JavascriptFiles() map[string]html5tag.Attributes {
	ret := make(map[string]html5tag.Attributes)
	p := path.Join(config.AssetPrefix, "messenger", "js", "goradd-sse.js")
	ret[p] = nil
	return ret
}

func (m *SseMessenger) Send(channel string, message string) {
	if m.hub != nil {
		m.hub.send <- clientMessage{channel, message}
	}
}

// MessageHandler handles incoming requests from clients.
//
// This handler should be registered at a convenient endpoint and that matches the javascript specified endpoint.
// It gets the client id from the context in the request. You should intercept
// the request, authorize the client, then insert the client ID into the context of the
// Request
func (m *SseMessenger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}
	clientID := r.FormValue("id")
	cmd := r.FormValue("cmd")

	switch cmd {
	case "new":
		_, ok := w.(http.Flusher)

		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		newClient(m.hub, w, r, clientID)

	case "subscribe":
		// subscription data is a list of channels
		var channels []string

		err := http2.ParseJsonBody(w, r, 10000, &channels)
		if err != nil {
			return
		}

		for _, channel := range channels {
			s := subscription{
				clientID: clientID,
				channel:  channel,
			}
			m.hub.subscribe <- s
		}
	}

}
