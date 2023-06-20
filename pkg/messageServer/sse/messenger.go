package sse

import (
	"fmt"
	"github.com/goradd/goradd/pkg/config"
	"github.com/goradd/goradd/pkg/goradd"
	http2 "github.com/goradd/goradd/pkg/http"
	_ "github.com/goradd/goradd/pkg/messageServer/ws/assets"
	"github.com/goradd/html5tag"
	"github.com/r3labs/sse"
	"net/http"
	"path"
)

type SseMessenger struct {
	Server *sse.Server
}

func (m *SseMessenger) Start() {
	m.Server = sse.New()
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
	m.Server.Publish(channel, &sse.Event{
		Data: []byte(message),
	})
}

// SocketHandler handles web socket requests to send messages to clients.
// It gets the client id from the context in the request. You should intercept
// the request, authorize the client, then insert the client ID into the context of the
// Request
func (m *SseMessenger) SocketHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.Context().Value(goradd.MessengerContext).(string)
		serveWs(m.hub, w, r, clientID)
	})
}
