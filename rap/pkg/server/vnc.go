package server

import (
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"sync"

	"github.com/elazarl/go-bindata-assetfs"

	"golang.org/x/net/context"
	"golang.org/x/net/websocket"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/ui/data/novnc"
	"github.com/tangfeixiong/go-to-kubernetes/rap/pb"
	"github.com/tangfeixiong/go-to-kubernetes/rap/pkg/gowebsockifynovnc"
)

func (ctl *controller) serve_noVNC(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := gowebsockifynovnc.FileServer(&assetfs.AssetFS{
		Asset:    novnc.Asset,
		AssetDir: novnc.AssetDir,
		Prefix:   "Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC",
	})
	prefix := "/novnc/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	wsmux := websocket.Server{
		Handshake: ctl.bootHandshake,
		Handler:   ctl.handleWss,
	}
	mux.Handle("/websockify", wsmux)

	ctl.wsnovnc = fileServer
}

func (ctl *controller) WebsockifyVNC(ctx context.Context, req *pb.VncReqResp) (*pb.VncReqResp, error) {
	fmt.Println("Requesting:", req)
	return ctl.wsnovnc.Dispatch(req)
}

func (ctl *controller) handleWss(wsconn *websocket.Conn) {
	fmt.Printf("handle wss: %+v\n", wsconn.Config())
	var cookie *http.Cookie
	var err error
	cookie, err = wsconn.Request().Cookie("novnc_token")
	if cookie == nil || err != nil {
		fmt.Println(wsconn.Request().URL, wsconn.Request().Form, wsconn.Request().Header)
		return
	}
	token := cookie.Value
	if token == "" {
		fmt.Println("Unexpected, token does not exist")
		return
	}

	var conn net.Conn
	var wsc *websocket.Conn
	var targetAddr string

	conn, wsc, targetAddr = ctl.wsnovnc.ValidateVNC(token)
	if targetAddr == "" {
		fmt.Println("API not invoked")
		return
	}
	if wsc != nil {
		if wsc == wsconn {
			fmt.Println("Already connected")
			return
		}
		fmt.Print("Disconnect old streaming")
		if err := wsc.Close(); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println()
		}
		if conn != nil {
			if err := conn.Close(); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println()
			}
		}
	}
	if conn == nil {
		fmt.Println("Connecting", targetAddr)
		conn, err = net.Dial("tcp", targetAddr)
		if err != nil {
			fmt.Println("Could not connect VNC, error:", err.Error())
			wsconn.Close()
			return
		}
	}
	ctl.wsnovnc.Connection(token, conn, wsconn)

	wsconn.PayloadType = websocket.BinaryFrame
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer wsconn.Close()
		l, e := io.Copy(conn, wsconn)
		fmt.Println("Client streaming terminated (ws -> vnc), ", l, e)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer conn.Close()
		l, e := io.Copy(wsconn, conn)
		fmt.Println("Server streaming terminated (vnc -> ws), ", l, e)
	}()
	//select {}
	wg.Wait()
	ctl.wsnovnc.Disconnection(token)
}

func (ctl *controller) bootHandshake(config *websocket.Config, r *http.Request) error {
	fmt.Printf("handshake: %+v\nrequest: %+v\n", config, r)
	config.Protocol = []string{"binary"}

	r.Header.Set("Access-Control-Allow-Origin", "*")
	r.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	return nil
}
