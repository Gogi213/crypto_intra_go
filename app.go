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

func StartServer(dataChannels [][]chan []byte, ports []string) {
	log.Println("Server started")

	fs := &fasthttp.FS{
		Root:       "templates",
		IndexNames: []string{"home.html"},
	}
	fsHandler := fs.NewRequestHandler()

	portIndex := 0
	for _, chs := range dataChannels {
		for _, ch := range chs {
			if portIndex >= len(ports) {
				log.Fatal("Not enough ports provided")
			}
			go startTCPListener(ports[portIndex], ch)
			portIndex++
		}
	}

	server := fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/":
				fsHandler(ctx)
			case "/ws1":
				handleWebsocket(ctx, dataChannels[0][0])
			case "/ws2":
				handleWebsocket(ctx, dataChannels[1][0])
			case "/ws3":
				handleWebsocket(ctx, dataChannels[2][0])
			case "/ws4":
				handleWebsocket(ctx, dataChannels[3][0])
			case "/ws5":
				handleWebsocket(ctx, dataChannels[4][0])
			case "/ws6":
				handleWebsocket(ctx, dataChannels[5][0])
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		},
	}

	log.Fatal(server.ListenAndServe(":8080"))
}

func startTCPListener(port string, dataChannel chan []byte) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn, dataChannel)
	}
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