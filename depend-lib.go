package main

import (
	"net/http"

	"github.com/actcat/sideci_go_sandbox/lib"
	"github.com/naoina/denco"
)

func HttpHandler(w http.ResponseWriter, r *http.Request, p denco.Params) {
	foo := p.Get("foo")
	w.Write([]byte(foo))
}

func UseLibFoo() {
	libfoo.Foo()
}
