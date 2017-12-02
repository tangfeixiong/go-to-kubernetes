package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var (
	sourceAddr = flag.String("source", "", "source address")
	targeAddr  = flag.String("target", "", "target address")
)

func main() {
	flag.Parse()
	if *sourceAddr == "" || *targeAddr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	mux := websocket.Server{
		Handshake: bootHandshake,
		Handler:   handleWss}

	http.Handle("/websockify", mux)
	err := http.ListenAndServe(*sourceAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleWss(wsconn *websocket.Conn) {
	log.Println("handlewss")
	conn, err := net.Dial("tcp", *targeAddr)

	if err != nil {
		log.Println(err)
		wsconn.Close()

	} else {
		wsconn.PayloadType = websocket.BinaryFrame

		go io.Copy(conn, wsconn)
		go io.Copy(wsconn, conn)

		select {}
	}
}

func bootHandshake(config *websocket.Config, r *http.Request) error {
	config.Protocol = []string{"binary"}

	r.Header.Set("Access-Control-Allow-Origin", "*")
	r.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	return nil
}
