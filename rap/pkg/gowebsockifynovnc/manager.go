package gowebsockifynovnc

import (
	"fmt"
	//"io"
	//"mime"
	"net"
	"net/http"
	//"net/textproto"
	//"net/url"
	//"os"
	"path"
	//"path/filepath"
	//"regexp"
	//"sort"
	"strings"
	//"time"

	"github.com/pborman/uuid"

	"golang.org/x/net/websocket"

	"github.com/tangfeixiong/go-to-kubernetes/rap/pb"
)

type scheduler struct {
	rr     *pb.VncReqResp
	conn   net.Conn
	wsconn *websocket.Conn
}

type WsnovncManager struct {
	root      http.FileSystem
	schedules map[string]*scheduler
	tokens    map[string]string
}

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
func FileServer(root http.FileSystem) /*http.Handler*/ *WsnovncManager {
	return &WsnovncManager{
		root:      root,
		schedules: make(map[string]*scheduler),
		tokens:    make(map[string]string),
	}
}

func (f *WsnovncManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if token := r.URL.Query().Get("token"); token != "" {
		fmt.Printf("request: %+v\n", r)
		http.SetCookie(w, &http.Cookie{
			Name:  "novnc_token",
			Value: token,
			Path:  "/websockify",
		})
	} else {
		fmt.Printf("header referer: %+v\n", r.Header.Get("Referer"))
	}
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	serveFile(w, r, f.root, path.Clean(upath), true)
}

func (mg *WsnovncManager) Dispatch(req *pb.VncReqResp) (*pb.VncReqResp, error) {
	resp := new(pb.VncReqResp)
	if req.VncAddr == "" {
		resp.StateCode = 10
		resp.StateMessage = "VNC Server Address required"
		return resp, fmt.Errorf(resp.StateMessage)
	}
	resp.VncAddr = req.VncAddr
	if req.AuthToken == "" {
		resp.StateCode = 11
		resp.StateMessage = "User token required"
		return resp, fmt.Errorf(resp.StateMessage)
	}
	resp.AuthToken = req.AuthToken

	resp.Token = uuid.New()
	mg.schedules[resp.Token] = &scheduler{
		rr: resp,
	}

	return resp, nil
}

func (mg *WsnovncManager) ValidateVNC(token string) (net.Conn, *websocket.Conn, string) {
	if _, ok := mg.schedules[token]; !ok {
		return nil, nil, ""
	}
	if mg.schedules[token].rr == nil {
		return nil, nil, ""
	}
	return mg.schedules[token].conn, mg.schedules[token].wsconn, mg.schedules[token].rr.VncAddr
}

func (mg *WsnovncManager) Token(id string) string {
	return mg.tokens[id]
}

func (mg *WsnovncManager) Connection(token string, conn net.Conn, wsconn *websocket.Conn) {
	mg.schedules[token].conn = conn
	mg.schedules[token].wsconn = wsconn
}
func (mg *WsnovncManager) Disconnection(token string) {
	//delete(mg.schedules, token)
	mg.schedules[token].conn = nil
	mg.schedules[token].wsconn = nil
}
