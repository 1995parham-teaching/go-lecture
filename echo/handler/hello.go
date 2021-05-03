package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbu-ce/ie-with-go/echo/request"
)

type Hello struct {
	From string
}

func (h Hello) User(c echo.Context) error {
	value := c.Param("username")

	log.Println(value)

	// nolint: wrapcheck
	return c.NoContent(http.StatusNoContent)
}

func (h Hello) Get(c echo.Context) error {
	if value := c.QueryParam("hello"); value != "" {
		log.Println(value)
	}

	// nolint: wrapcheck
	return c.JSON(http.StatusOK, fmt.Sprintf("Hello World from %s", h.From))
}

func (h Hello) Post(c echo.Context) error {
	var req request.Name

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if req.Count == nil {
		log.Println("There is no count")
	} else {
		log.Printf("There is a count %d", *req.Count)
	}

	// nolint: wrapcheck
	return c.String(http.StatusOK, fmt.Sprintf("Hello to %s from %s", req.Name, h.From))
}
