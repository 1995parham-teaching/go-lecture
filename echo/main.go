package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sbu-ce/ie-with-go/echo/handler"
)

func main() {
	e := echo.New()

	h := handler.Hello{
		From: "Golang",
	}

	e.GET("/hello", h.Get)

	e.POST("/hello", h.Post)

	e.GET("/hello/:username", h.User)

	e.Start("0.0.0.0:1373")
}
