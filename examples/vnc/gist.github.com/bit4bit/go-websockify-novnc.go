package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/philips/go-bindata-assetfs"

	"golang.org/x/net/websocket"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/ui/data/novnc"
)

var (
	sourceAddr = flag.String("source", "0.0.0.0:1987", "source address")
	targeAddr  = flag.String("target", "172.17.0.4:5900", "target address")
)

func main() {
	flag.Parse()
	if *sourceAddr == "" || *targeAddr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/ on <host>/
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    novnc.Asset,
		AssetDir: novnc.AssetDir,
		Prefix:   "Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC",
	})
	prefix := "/novnc/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request", r.URL.Path)
		if r.URL.Path == "/favicon.ico" {
			http.Redirect(w, r, "/novnc/vnc.html", http.StatusMovedPermanently)
		} else {
			mux.ServeHTTP(w, r)
		}
	}))

	wsmux := websocket.Server{
		Handshake: bootHandshake,
		Handler:   handleWss,
	}
	//http.Handle("/websockify", wsmux)
	mux.Handle("/websockify", wsmux)

	// err := http.ListenAndServe(*sourceAddr, nil)
	err := http.ListenAndServe(*sourceAddr, allowCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	fmt.Printf("preflight request for %s\n", r.URL.Path)
	return
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
