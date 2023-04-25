package http

import (
	"fmt"
	"net/http"
	"testing"

	ws "github.com/gorilla/websocket"

	"github.com/wrmsr/bane/core/check"
)

func runWsServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`
<!DOCTYPE html>
<html lang="en" data-layout="responsive">

<head>
	<script defer>
		var socket = new WebSocket("ws://localhost:8080/echo");

		socket.onopen = function () {
			var output = document.getElementById("output");
			output.innerHTML += "Status: Connected\n";
		};

		socket.onmessage = function (e) {
			var output = document.getElementById("output");
			output.innerHTML += "Server: " + e.data + "\n";
		};

		function send() {
			var input = document.getElementById("input");
			socket.send(input.value);
			input.value = "";
		}
	</script>
</head>

<body>
	<input id="input" type="text" />
	<button onclick="send()">Send</button>
	<pre id="output"></pre>
</body>

</html>
`,
		))
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		upgrader := ws.Upgrader{}
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	check.Must(http.ListenAndServe(":8080", nil))
}

func TestWebSockets(t *testing.T) {
	t.Skip("local only")

	runWsServer()
}
