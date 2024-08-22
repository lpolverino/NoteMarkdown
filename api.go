package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
	ListenAddr int
}

func CreateApi(listenAddr int) *Api {
	return &Api{ListenAddr: listenAddr}
}

func (a *Api) run() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/note", func(c echo.Context) error {

		contentType := c.Request().Header.Get("Content-Type")

		if contentType != "multipart/form-data" {
			return http.ErrNotMultipart
		}
		return c.String(http.StatusOK, "Note Saves")
	})

	e.POST("/grammar", func(c echo.Context) error {
		return c.String(http.StatusOK, "Grammar cheked")
	})

	e.GET("/note/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
