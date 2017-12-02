package gowebsockifynovnc

/*
  Inspired by
    https://github.com/golang/go/blob/master/src/net/http/sniff.go
*/

// The algorithm uses at most sniffLen bytes to make its decision.
const sniffLen = 512
