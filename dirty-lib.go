package main

import (
	"net/http"

	"github.com/actcat/sideci_go_sandbox/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/denco"
	_ "github.com/naoina/kocha"
	_ "github.com/revel/revel"
	_ "github.com/zenazn/goji"
)

func HttpHandler(w http.ResponseWriter, r *http.Request, p denco.Params) {
	foo := p.Get("foo")
	w.Write([]byte(foo))
}

func UseLibFoo() {
	libfoo.Foo()
}
