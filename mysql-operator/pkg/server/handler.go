package server

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"path"
	"strings"
	"sync"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/golang/glog"
	//"github.com/gorilla/websocket"

	"golang.org/x/net/context"
	"golang.org/x/net/websocket"

	"github.com/tangfeixiong/go-to-bigdata/nps-wss/pkg/httpfs"
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pb"
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/ui/data/webapp"
)

func (ctl *controller) CreateCrd(ctx context.Context, req *pb.CrdReqResp) (*pb.CrdReqResp, error) {
	fmt.Println("Request to create CRD:", req)
	resp := &pb.CrdReqResp{
		Recipe: new(pb.CrdRecipient),
	}
	if req == nil || req.Recipe == nil {
		resp.StateCode = 100
		resp.StateMessage = "CRD recipe is required"
		return resp, errors.New(resp.StateMessage)
	}
	if req.Recipe.Name == "" || req.Recipe.Group == "" {
		resp.StateCode = 101
		resp.StateMessage = "Empty CRD field is not allowed"
		return resp, errors.New(resp.StateMessage)
	}

	resp.Recipe.Name = req.Recipe.Name
	resp.Recipe.Group = req.Recipe.Group
	resp.Recipe.Version = req.Recipe.Version
	resp.Recipe.Scope = req.Recipe.Scope
	resp.Recipe.Plural = req.Recipe.Plural
	resp.Recipe.Singular = req.Recipe.Singular
	resp.Recipe.Kind = req.Recipe.Kind
	err := ctl.ops["mysql-operator"].CreateCRD(req.Recipe)
	if err != nil {
		glog.Infof("Create CRD failed: %s", err.Error())
		resp.StateCode = 10
		resp.StateMessage = err.Error()
		return resp, err
	}
	return resp, nil
}

func (ctl *controller) ReapCrd(ctx context.Context, req *pb.CrdReqResp) (*pb.CrdReqResp, error) {
	fmt.Println("Requesting:", req)
	return new(pb.CrdReqResp), nil
}

/*
  https://github.com/golang/go/blob/master/net/http/fs.go
*/
// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at root.
//
// To use the operating system's file system implementation,
// use http.Dir:
//
//     http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
// As a special case, the returned file server redirects any request
// ending in "/index.html" to the same path, without the final
// "index.html".
/*
func FileServer(root http.FileSystem) http.Handler {
	return &fileHandler{root}
}
*/

func (ctl *controller) serveWebPages(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	//	fileserver := FileServer(&assetfs.AssetFS{
	//		Asset:    webapp.Asset,
	//		AssetDir: webapp.AssetDir,
	//		Prefix:   "static",
	//	})
	//	prefix := "/html/"
	//	mux.Handle(prefix, http.StripPrefix(prefix, fileserver))
	ctl.root = &assetfs.AssetFS{
		Asset:    webapp.Asset,
		AssetDir: webapp.AssetDir,
		Prefix:   "static",
	}
	prefix := "/test/"
	mux.Handle(prefix, http.StripPrefix(prefix, ctl))

	//	ws := websocket.Server{
	//		Handshake: ctl.bootHandshake,
	//		Handler:   ctl.handleWss,
	//	}
	//	mux.Handle("/ws", ws)
	prefix += "ws"
	mux.HandleFunc(prefix, ctl.WebsocketHandler)
}

func (f *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if token := r.URL.Query().Get("token"); token != "" {
		fmt.Printf("request: %+v\n", r)
		http.SetCookie(w, &http.Cookie{
			Name:  "nps_token",
			Value: token,
			Path:  "/",
		})
	} else {
		fmt.Printf("header referer: %+v\n", r.Header.Get("Referer"))
	}

	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	httpfs.ServeFileSystem(w, r, f.root, path.Clean(upath), true)
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

	conn, wsc, targetAddr = nil, nil, "127.0.0.1:5670"
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
}

func (ctl *controller) bootHandshake(config *websocket.Config, r *http.Request) error {
	fmt.Printf("handshake: %+v\nrequest: %+v\n", config, r)
	config.Protocol = []string{"binary"}

	r.Header.Set("Access-Control-Allow-Origin", "*")
	r.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	return nil
}
