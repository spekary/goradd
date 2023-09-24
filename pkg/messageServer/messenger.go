// Package messageServer implements a messaging service that sends messages to the http client.
//
// There are many mechanisms for doing this, and new ones are evolving constantly.
// Goradd uses this service to communicate to the client that data has changed and it should initiate a
// data refresh. However, it is general purpose enough that the application or specific controls could make
// use of it as well to notify the client of anything.
//
// The original implementation used WebSockets and was based on the gorilla websocket library. That library
// has been deprecated, and the new default uses a server-side events library. However, MTTQ, NATS, PubNub
// and other implementations are possible.

package messageServer

import (
	"encoding/json"
	"github.com/goradd/goradd/pkg/log"
	"github.com/goradd/html5tag"
)

var Messenger MessengerI

type MessengerI interface {
	// JavascriptInit returns javascript that needs to be called when a page starts to initialize
	JavascriptInit() string
	// Send sends a message to the given channel
	Send(channel string, message string)
	// JavascriptFiles returns file names and attributes used to embed the javascript files needed by your messenger
	JavascriptFiles() map[string]html5tag.Attributes
}

// Send sends a message to the given channel via the current messenger
func Send(channel string, message interface{}) {
	msg, err := json.Marshal(message)
	if err != nil {
		log.Debug(err)
		return
	}
	if Messenger != nil {
		Messenger.Send(channel, string(msg))
	}
}
