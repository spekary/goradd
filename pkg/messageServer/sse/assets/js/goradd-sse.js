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
        goradd._messageLoc += window.location.hostname + port + loc;

        if (!goradd._sse) {
            goradd._sse = new EventSource(goradd._messageLoc + "?cmd=new&id=" + goradd.getPageState());

            goradd._sse.onopen = function() {
                console.log('SSE Open');
                goradd.subscribeWatchers();
                g$(goradd.form()).trigger("messengerReady");
            }
            goradd._sse.onerror= function(err) {
                console.error("EventSource failed:", err);
            }
            goradd._sse.onmessage = function(e) {
                goradd._handleSseMessage(e);
            }
        }

        window.onbeforeunload = function(e) {
            goradd._sse.close();
        }
    }
};

// channels is an array of strings indicating the channels to subscribe to
// f is the function to call when a message comes through on that channel
goradd.subscribe = function(channels, f) {
    goradd._sendMessage("subscribe", channels);
    goradd.each(channels, function() {
        goradd._channels[this] = f;
    });
};

goradd._sendMessage = function(cmd, data) {
    let params = new URLSearchParams();
    let options = {};

    params.append('id', goradd.getPageState());
    params.append('cmd', cmd);

    if (data) {
        options = {
            method: 'POST',
            body: JSON.stringify( data ),
            headers: {
                "Content-Type": "application/json",
            },
            cache: 'no-store',
        };
    }
    fetch(goradd._messageLoc + "?" + params.toString(), options);
}

/*
 *   The default message handler. Will route the message to the appropriate channel.
 */
goradd._handleSseMessage = function(e) {
    console.log("messages: " + e.data);
    let messages = JSON.parse(e.data);

    goradd.each(messages, function() {
        let msg = this;
        let f = goradd._channels[msg.channel];
        if (!!f) {
            f(msg);
        }
    });
};


