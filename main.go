package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/sideci/go_private_library"
)

func main() {
	e := echo.New()
	fmt.Println(go_private_library.Fibo(10))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Run(standard.New(":1323"))
}
