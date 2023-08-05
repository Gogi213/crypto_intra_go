// app.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  7024,
	WriteBufferSize: 7024,
}

func StartGin(dataChannel1, dataChannel2 chan []byte) {
	r := gin.Default()
	r.LoadHTMLFiles("templates/home.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		go handleConnection(conn, dataChannel1)
		go handleConnection(conn, dataChannel2)
	})

	r.Run()
}

func handleConnection(conn *websocket.Conn, dataChannel chan []byte) {
	for {
		message := <-dataChannel
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
}
