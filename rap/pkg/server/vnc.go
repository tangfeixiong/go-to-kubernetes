package server

import (
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"

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
	cookie, err = wsconn.Request().Cookie("JSESSIONID")
	if cookie == nil || err != nil {
		fmt.Println(wsconn.Request().URL, wsconn.Request().Form, wsconn.Request().Header)
		return
	}
	token := ctl.wsnovnc.Token(cookie.Value)
	if token == "" {
		fmt.Println("Unexpected, token does not exist")
		return
	}

	var conn net.Conn
	var targetAddr string

	conn, targetAddr = ctl.wsnovnc.ValidateVNC(token)
	if targetAddr == "" {
		fmt.Println("API not invoked")
		return
	}
	if conn == nil {
		fmt.Println("First connecting")
		conn, err = net.Dial("tcp", targetAddr)
		if err != nil {
			fmt.Println("Could not connect VNC, error:", err.Error())
			wsconn.Close()
			return
		}
		ctl.wsnovnc.Connection(token, conn)
	}

	wsconn.PayloadType = websocket.BinaryFrame

	go io.Copy(conn, wsconn)
	go io.Copy(wsconn, conn)

	select {}

}

func (ctl *controller) bootHandshake(config *websocket.Config, r *http.Request) error {
	fmt.Println(r.URL, r.Form, r.Header)
	config.Protocol = []string{"binary"}

	r.Header.Set("Access-Control-Allow-Origin", "*")
	r.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	return nil
}
