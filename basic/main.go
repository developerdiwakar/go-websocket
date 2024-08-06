package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write TEXT
		// w.Write([]byte("WebSocket"))

		// Write HTML content
		// 	htmlContent := `
		// <!DOCTYPE html>
		// <html>
		// <head>
		// 	<title>Go Web Server</title>
		// </head>
		// <body>
		// 	<h1>Welcome to my Go Web Server</h1>
		// 	<p>This is a simple example of serving HTML content.</p>
		// </body>
		// </html>
		// `
		// 	w.Header().Set("Content-type", "text/html")
		// 	fmt.Fprint(w, htmlContent)

		// Write HTML File
		http.ServeFile(w, r, "./basic/index.html")
	})

	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		log.Printf("Received: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
