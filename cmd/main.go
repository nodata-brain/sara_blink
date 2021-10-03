package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", root)

	e.Logger.Fatal(e.Start(":8080"))
}

func root(c echo.Context) error {

}
