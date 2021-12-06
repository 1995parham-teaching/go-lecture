package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/cng-by-example/hello-go/echo/handler"
)

func main() {
	e := echo.New()

	h := handler.Hello{
		From: "Golang",
	}

	e.GET("/hello", h.Get)

	e.POST("/hello", h.Post)

	e.GET("/hello/:username", h.User)

	if err := e.Start("0.0.0.0:1373"); err != nil {
		log.Fatal(err)
	}
}
