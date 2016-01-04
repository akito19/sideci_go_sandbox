package main

import (
	"net/http"

	"github.com/naoina/denco"
)

func HttpHandler(w http.ResponseWriter, r *http.Request, p denco.Params) {
	foo := p.Get("foo")
	w.Write([]byte(foo))
}
