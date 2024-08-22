package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
		fmt.Println(contentType)

		if !strings.Contains(contentType, "multipart/form-data") {
			fmt.Println("here")
			return http.ErrNotMultipart
		}

		file, err := c.FormFile("file")
		if err != nil {
			fmt.Printf("cannot load file, %v\n", err)
			return err
		}

		src, err := file.Open()
		if err != nil {
			fmt.Printf("cannot open file, %v\n", err)
			return err
		}
		defer src.Close()

		dst, err := os.Create(file.Filename)
		if err != nil {
			fmt.Printf("cannot create file, %v\n", err)
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			fmt.Printf("cannot copy the file, %v", err)
			return err
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
