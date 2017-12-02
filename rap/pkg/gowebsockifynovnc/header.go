package gowebsockifynovnc

import (
	"net/http"
)

/*
  Inspired by
    https://github.com/golang/go/blob/master/src/net/http/header.go
*/

// get is like Get, but key must already be in CanonicalHeaderKey form.
//func (h Header) get(key string) string {
func header_get(h http.Header, key string) string {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}
