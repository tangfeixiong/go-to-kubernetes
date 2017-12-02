# Help

Inspired by:

* https://github.com/mitchellh/go-vnc
* https://github.com/openai/go-vncdriver
* https://github.com/bradfitz/rfbgo

## Encodings

x11vnc

according http://karlrunge.com/x11vnc/x11vnc_opts.html
```
vncviewer -encodings 'copyrect tight zrle hextile' localhost:0
```

## go-vnc

example 1
```
[vagrant@localhost go-to-kubernetes]$ CGO_ENABLED=0 go run  --installsuffix cgo  -tags no_gl ./examples/vnc/mitchellh0x2Fgo-vnc/main.go -addr=172.17.0.4:5900
connected
encoding sent
frame buffer update requested
2017/11/29 19:00:54 Just received: *vnc.FramebufferUpdateMessage &{Rectangles:[{X:0 Y:0 Width:1024 Height:768 Enc:0xc420082040}]}
1
786432
1024 768
writed into bmp
```

## noVNC

generate go code
```
[vagrant@localhost go-to-kubernetes]$ go-bindata -nocompress -o pkg/ui/data/novnc/datafile.go -prefix `pwd` -pkg novnc /Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/app/... /Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/core/... /Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/LICENSE.txt /Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/vendor/... /Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/vnc.html /Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/vnc_lite.html 
```

```
[vagrant@localhost go-to-kubernetes]$ go-bindata -nocompress -o pkg/ui/data/novnc/datafile.go -prefix `pwd` -pkg novnc -ignore "config|description|FETCH_HEAD|HEAD|index|ORIG_HEAD|packed-refs|shallow|\\.sample|exclude|master|\\.idx|\\.pack|\\.gitmodules|v[01]\\.[0-6]|\\.gitignore|\\.npmignore|\\.travis" /Users/fanhongling/Downloads/workspace/src/github.com/novnc/noVNC/...
```