package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"runtime"
	// "strconv"

	vncclient "github.com/mitchellh/go-vnc"
)

type config struct {
	address  string
	password string
	encoding string
}

var vnc_config *config

func init() {
	vnc_config = &config{
		encoding: "tight",
		password: "",
	}

	flag.StringVar(&vnc_config.address, "addr", "", "vnc server")
	flag.StringVar(&vnc_config.password, "pass", "", "vnc password")
	flag.StringVar(&vnc_config.encoding, "enc", "", "vnc encoding: none or raw")
}

func main() {
	flag.Parse()
	nc, err := net.Dial("tcp", vnc_config.address)
	if err != nil {
		log.Fatalf("error connecting to mock server: %s", err)
	}

	target := nc
	var conn *vncclient.ClientConn
	serverMessageCh := make(chan vncclient.ServerMessage)
	if vnc_config.password == "" {
		conn, err = vncclient.Client(target, &vncclient.ClientConfig{
			Auth: []vncclient.ClientAuth{
				new(vncclient.ClientAuthNone),
			},
			ServerMessageCh: serverMessageCh,
		})
	} else {
		conn, err = vncclient.Client(target, &vncclient.ClientConfig{
			Auth: []vncclient.ClientAuth{
				&vncclient.PasswordAuth{
					Password: vnc_config.password,
				},
			},
			ServerMessageCh: serverMessageCh,
		})
	}
	if err != nil {
		log.Fatalf("conn err: %v", err)
	}
	if conn == nil {
		log.Fatal("conn Unexpected")
	}
	println("connected")

	enc := new(vncclient.RawEncoding)
	encodings := []vncclient.Encoding{enc}
	err = conn.SetEncodings(encodings)
	if err != nil {
		log.Fatalf("could not set encodings: %v", err)
	}
	println("encoding sent")

	//					rect := vncclient.Rectangle{
	//						X:      0,
	//						Y:      0,
	//						Width:  64,
	//						Height: 48,
	//						Enc:    enc,
	//					}
	// var codec vncclient.Encoding = enc
	w := uint16(1024)
	h := uint16(768)
	//	w = uint16(160)
	//	h = uint16(90)
	err = conn.FramebufferUpdateRequest(true, 0, 0, w, h)
	if err != nil {
		log.Fatalf("could not set FramebufferUpdateRequest: %v", err)
	}
	println("frame buffer update requested")

	for {
		select {
		case msg := <-serverMessageCh:
			log.Printf("Just received: %T %+v\n", msg, msg)
			switch msg := msg.(type) {
			case *vncclient.FramebufferUpdateMessage:
				//				updates <- msg
				//				// Keep re-requesting!
				//				c.updated.L.Lock()
				//				if !c.pauseUpdates {
				//					err := c.requestUpdate()
				//					if err != nil {
				//						select {
				//						case c.mgr.Error <- err:
				//						default:
				//						}
				//					}
				//				}
				//				c.updated.L.Unlock()
				update := (*vncclient.FramebufferUpdateMessage)(msg)
				println(len(update.Rectangles))
				/*
					fanhonglingdeMacBook-Pro:~ fanhongling$ hexdump test.bmp
					0000000 42 4d f6 00 00 00 00 00 00 00 36 00 00 00 28 00
					0000010 00 00 08 00 00 00 08 00 00 00 01 00 18 00 00 00
					0000020 00 00 c0 00 00 00 00 00 00 00 00 00 00 00 00 00
					0000030 00 00 00 00 00 00 ff ff ff ff ff ff ff ff ff ff
					0000040 ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff
					*
					00000f0 ff ff ff ff ff ff
				*/
				for _, rect := range update.Rectangles {
					switch enc := rect.Enc.(type) {
					case *vncclient.RawEncoding:
						println(len(enc.Colors))
						buf := &bytes.Buffer{}
						//The characters "BM"
						buf.Write([]byte{'B', 'M'})
						//The size of the file in bytes
						var size uint32 = 14 + 40 + uint32(w)*uint32(h)*3
						buf.Write([]byte{uint8(size & 0xff), uint8(size >> 8 & 0xff), uint8(size >> 16 & 0xff), uint8(size >> 24 & 0xff)})
						//Unused - must be zero
						buf.Write([]byte{0, 0, 0, 0})
						//Offset to start of Pixel Data
						buf.Write([]byte{54, 0, 0, 0})
						//Header Size - Must be at least 40
						buf.Write([]byte{40, 0, 0, 0})
						//Image width in pixels
						buf.Write([]byte{uint8(w & 0xff), uint8(w >> 8 & 0xff), 0, 0})
						//Image height in pixels
						var neg uint32 = 0xffffffff - uint32(h) + 1
						buf.Write([]byte{uint8(neg & 0xff), uint8(neg >> 8 & 0xff), uint8(neg >> 16 & 0xff), uint8(neg >> 24 & 0xff)})
						//buf.Write([]byte{uint8(h & 0xff), uint8(h >> 8 & 0xff), 0, 0})
						//Must be 1
						buf.Write([]byte{1, 0})
						//Bits per pixel - 1, 4, 8, 16, 24, or 32
						buf.Write([]byte{24, 0})
						//Compression type (0 = uncompressed)
						buf.Write([]byte{0, 0, 0, 0})
						//Image Size - may be zero for uncompressed images
						buf.Write([]byte{0, 0, 0, 0})
						//Preferred resolution in pixels per meter
						buf.Write([]byte{0, 0, 0, 0})
						//Preferred resolution in pixels per meter
						buf.Write([]byte{0, 0, 0, 0})
						//Number Color Map entries that are actually used
						buf.Write([]byte{0, 0, 0, 0})
						//Number of significant colors
						buf.Write([]byte{0, 0, 0, 0})
						var n, l int
						for _, colour := range enc.Colors {
							//							if l%w == 0 {
							//								buf.Write([]byte{0, 0, 0, 0})
							//							}
							l++
							if colour.R >= 255 || colour.G >= 255 || colour.B >= 255 {
								fmt.Println(l, colour.R, colour.G, colour.B)
							}
							n, err = buf.Write([]byte{uint8(colour.R * 4 & 0xff), uint8(colour.G * 3 & 0xff), uint8(colour.B * 5 & 0xff)})
							if err != nil {
								log.Fatalf("could not read : %v", err)
							}
							if n == 0 {
								log.Fatalf("Read unexpected")
							}
						}
						fmt.Println(w, l/int(w))
						err = ioutil.WriteFile("temp.bmp", buf.Bytes(), 0644)
						if err != nil {
							log.Fatalf("could not write file : %v", err)
						}
						println("writed into bmp")
						return
					default:
						log.Fatalf("unsupported encoding: %T", enc)
					}
				}
				//		case <-c.mgr.Done:
				//			log.Debugf("[%s] server message goroutine exiting", c.label)
				//			return nil
			}
		}
	}

	runtime.Goexit()
}
