package server

import (
	"fmt"
	"mime"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/grpc-gateway-example/pkg/ui/data/swagger"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pb"
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/operator"
)

type controller struct {
	config *Config
	ops    map[string]*operator.Operator
	root   http.FileSystem
	signal os.Signal
}

func Start(config *Config) {
	ctl := &controller{
		config: config,
		ops:    make(map[string]*operator.Operator),
	}
	op, err := operator.Run(config.PrebootCfg)
	if err != nil {
		glog.Errorf("Start operator failed: %v", err)
		return
	}
	ctl.ops["mysql-operator"] = op
	ctl.start()
}

func (ctl *controller) start() {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctl.startGRPC(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctl.startGateway(ch)
	}()

	/*
	   https://github.com/kubernetes/kubernetes/blob/release-1.1/build/pause/pause.go
	*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	// Block until a signal is received.
	ctl.signal = <-c

	//wg.Wait()
}

func (ctl *controller) startGRPC(ch chan<- string) {
	s := grpc.NewServer()

	pb.RegisterSimpleGRpcServiceServer(s, ctl)
	host := ctl.config.SecureAddress

	l, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start gRPC into", l.Addr())
	go func() {
		time.Sleep(500)
		ch <- host
	}()
	if err := s.Serve(l); nil != err {
		panic(err)
	}
}

func (ctl *controller) startGateway(ch <-chan string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	// mux.HandleFunc("/swagger/", serveSwagger2)
	//	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	//		io.Copy(w, strings.NewReader(healthcheckerpb.Swagger))
	//	})

	dopts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	gRPCHost := <-ch

	host := ctl.config.InsecureAddress
	if err := pb.RegisterSimpleGRpcServiceHandlerFromEndpoint(ctx, gwmux, gRPCHost, dopts); err != nil {
		fmt.Println("Register gRPC gateway failed.", err)
		return
	}

	mux.Handle("/", gwmux)
	// serveSwagger(mux)
	ctl.serveWebPages(mux)

	lstn, err := net.Listen("tcp", host)
	if nil != err {
		panic(err)
	}

	fmt.Println("Start gRPC Gateway into", lstn.Addr())
	//	if err := http.ListenAndServe(host, allowCORS(mux)); nil != err {
	//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
	//	}
	s := &http.Server{
		Handler: func /*allowCORS*/ (h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if origin := r.Header.Get("Origin"); origin != "" {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
						func /*preflightHandler*/ (w http.ResponseWriter, r *http.Request) {
							headers := []string{"Content-Type", "Accept"}
							w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
							methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
							w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
							glog.Infof("preflight request for %s", r.URL.Path)
							return
						}(w, r)
						return
					}
				}
				h.ServeHTTP(w, r)
			})
		}(mux),
	}

	if err := s.Serve(lstn); nil != err {
		fmt.Fprintln(os.Stderr, "Server died.", err.Error())
	}
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
