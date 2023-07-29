// app.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
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

		connRabbit, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Fatal("Failed to connect to RabbitMQ: ", err)
		}
		defer connRabbit.Close()

		ch, err := connRabbit.Channel()
		if err != nil {
			log.Fatal("Failed to open a channel: ", err)
		}
		defer ch.Close()

		msgs, err := ch.Consume(
			"binance_queue",
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatal("Failed to register a consumer: ", err)
		}

		for d := range msgs {
			if err := conn.WriteMessage(websocket.TextMessage, d.Body); err != nil {
				return
			}
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
