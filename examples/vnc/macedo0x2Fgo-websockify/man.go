package main

import (
	"flag"
	"log"
	"mime"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/philips/go-bindata-assetfs"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/ui/data/novnc"
)

var (
	bind_addr   string
	vnc_addr    string
	logger      *log.Logger = log.New(os.Stdout, "[go-websockify] ", log.Lshortfile|log.LstdFlags)
	proxyserver *ProxyServer
)

func init() {
	flag.StringVar(&bind_addr, "bind-addr", ":1987", "The address this proxy used, default 0.0.0.0:1987")
	flag.StringVar(&vnc_addr, "vnc-addr", "", "The VNC server this proxy connect to, for example 172.17.0.4:5900")
}

func main() {
	flag.Parse()
	logger.Printf("listening on %s\n", bind_addr)

	mux := http.NewServeMux()

	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC// on <host>/
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    novnc.Asset,
		AssetDir: novnc.AssetDir,
		Prefix:   "Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC",
	})
	prefix := "/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	//http.HandleFunc("/", handleConnection)
	mux.HandleFunc("/", handleConnection)

	if err := http.ListenAndServe(bind_addr, allowCORS(mux)); err != nil {
		panic(err)
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	logger.Printf("new connection from: %s", r.RemoteAddr)
	wsConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		logger.Fatal(err)
		return
	}

	dataType, data, err := wsConn.ReadMessage()
	if err != nil {
		_ = wsConn.WriteMessage(dataType, []byte("Fail: "+err.Error()))
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", string(data))
	if err != nil {
		errorMsg := "FAIL(net resolve tcp addr): " + err.Error()
		logger.Println(errorMsg)
		_ = wsConn.WriteMessage(websocket.CloseMessage, []byte(errorMsg))
		return
	}

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		errorMsg := "FAIL(net dial tcp): " + err.Error()
		logger.Println(errorMsg)
		_ = wsConn.WriteMessage(websocket.CloseMessage, []byte(errorMsg))
		return
	}

	proxyserver = NewProxyServer(wsConn, tcpConn)
	go proxyserver.doProxy()
}

type ProxyServer struct {
	wsConn  *websocket.Conn
	tcpConn *net.TCPConn
}

func NewProxyServer(wsConn *websocket.Conn, tcpConn *net.TCPConn) *ProxyServer {
	proxyserver := ProxyServer{wsConn, tcpConn}
	return &proxyserver
}

func (proxyserver *ProxyServer) doProxy() {
	go proxyserver.wsToTcp()
	proxyserver.tcpToWs()
}

func (proxyserver *ProxyServer) tcpToWs() {
	buffer := make([]byte, 1024)

	for {
		n, err := proxyserver.tcpConn.Read(buffer)
		if err != nil {
			proxyserver.tcpConn.Close()
			break
		}

		err = proxyserver.wsConn.WriteMessage(websocket.BinaryMessage, buffer[0:n])
		if err != nil {
			logger.Println(err.Error())
		}
	}
}

func (proxyserver *ProxyServer) wsToTcp() {
	for {
		_, data, err := proxyserver.wsConn.ReadMessage()
		if err != nil {
			break
		}

		_, err = proxyserver.tcpConn.Write(data)
		if err != nil {
			logger.Println(err.Error())
			break
		}
	}
}
