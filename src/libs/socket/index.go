package socketIO

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

var SocketServer *socketio.Server

// InitSocket handler for use by app
func InitSocket() error {
	SocketServer = socketio.NewServer(nil)
	return nil
}

// SocketEvents from websocket
func SocketEvents() {
	SocketServer.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Println("socket connected:", conn.ID())
		return nil
	})

	SocketServer.OnEvent("/", "bye", func(s socketio.Conn, msg string) string {
		fmt.Println(msg)
		log.Println(s.Close())
		return msg
	})

	SocketServer.OnError("/", func(conn socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})

	SocketServer.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		fmt.Println("closed:", reason)
	})
}
