package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HelloWorld (c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// TODO: IMPLEMENT PART A