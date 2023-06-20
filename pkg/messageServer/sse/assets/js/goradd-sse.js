/*

Goradd Server-Sent Event Client

This file attaches a server-sent event client to the current goradd form. The client replies to one message,
which will cause the form to do an ajax update.

The point of all this is to allow the server to simulate a push update of the form. Essentially,
the server is just telling the client that its time to pull updates.

To implement this differently, you would need to substitute this js with a system that subscribes
to a channel using the pagestate as the channel id. On the server end, you would send to the channel
the update message. Just replace the initMessageClient function with your own.
*/

goradd._channels = {};
goradd._messageLoc = "";

goradd.initMessagingClient = function(loc) {
    if (typeof(EventSource) !== "undefined") {
        var port;

        if (location.protocol === 'https:') {
            goradd._messageLoc = "https://";
        } else {
            goradd._messageLoc = "http://";
        }
        port = location.port ? ":" + location.port : "";
        goradd._messageLoc += window.location.hostname + port + loc + "?id=" + goradd.getPageState();
        if (!goradd._sse) {
            goradd._sse = new EventSource(goradd._messageLoc);
        }
        goradd._sse.addEventListener("message", goradd._handleSseMessage);
        // we purposefully do not use goradd._sse.onmessage = ... so that we can add multiple event listeners.
        goradd._sse.onopen = function() {
            goradd.subscribeWatchers();
            g$(goradd.form()).trigger("messengerReady");
        }
    }
};

// channels is an array of strings indicating the channels to subscribe to
// f is the function to call when a message comes through on that channel
goradd.subscribe = function(channels, f) {
    var msg = {};
    msg["subscribe"] = channels;
    fetch(goradd._messageLoc, );
    goradd._sse.send(JSON.stringify(msg));
    goradd.each(channels, function() {
        goradd._channels[this] = f;
    });
};

/*
The default message handler. Will route the message to the appropriate channel.
 */
goradd._handleSseMessage = function(e) {
    var messages = JSON.parse(e.data);
    console.log("messages: " + e.data);

    goradd.each(messages, function() {
        var msg = this;
        var f = goradd._channels[msg.channel];
        if (!!f) {
            f(msg);
        }
    });
};

/*
Close the websocket connection.
See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/close for status codes
*/
goradd._closeWebSocket = function(status) {
    if (goradd._ws) {
        goradd._ws.close(); // Not all browsers support a status code, and attempting one breaks javascript.
    }
};

goradd._handleWsClose = function(e) {
    goradd._channels = {};
    goradd._ws = null;
};

