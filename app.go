package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartGin() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		// Here you should replace with your own data source
		// For example, you can use a channel to receive data
		dataChannel := make(chan []byte)

		go func() {
			for dataBytes := range dataChannel {
				var data Message
				if err := proto.Unmarshal(dataBytes, &data); err != nil {
					log.Fatal("Failed to unmarshal data: ", err)
				}
				if err := conn.WriteMessage(websocket.TextMessage, dataBytes); err != nil {
					return
				}
			}
		}()
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
