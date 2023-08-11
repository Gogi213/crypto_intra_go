// app.go
// app.go
package main

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net"
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  200000,
	WriteBufferSize: 200000,
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
}

func StartServer(dataChannels []chan []byte, port1 string, port2 string) {
	log.Println("Server started")

	fs := &fasthttp.FS{
		Root:       "templates",
		IndexNames: []string{"home.html"},
	}
	fsHandler := fs.NewRequestHandler()

	go func() {
		ln, err := net.Listen("tcp", ":"+port1)
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleConn(conn, dataChannels[0])
		}
	}()

	go func() {
		ln, err := net.Listen("tcp", ":"+port2)
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleConn(conn, dataChannels[1])
		}
	}()

	server := fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/":
				fsHandler(ctx)
			case "/ws1":
				handleWebsocket(ctx, dataChannels[0])
			case "/ws2":
				handleWebsocket(ctx, dataChannels[1])
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		},
	}

	log.Fatal(server.ListenAndServe(":8080"))
}

func handleWebsocket(ctx *fasthttp.RequestCtx, dataChannel chan []byte) {
	err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		defer conn.Close()

		for data := range dataChannel {
			err := conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err)
				break
			}
		}
	})

	if err != nil {
		log.Println(err)
	}
}

func handleConn(conn net.Conn, dataChannel chan []byte) {
	defer conn.Close()
	for {
		data := make([]byte, 20000)
		n, err := conn.Read(data)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed")
				return
			}
			log.Println(err)
			return
		}
		dataChannel <- data[:n]
	}
}