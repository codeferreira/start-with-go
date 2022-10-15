package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	server := echo.New()

	server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello World!")
	})

	server.Logger.Fatal(server.Start(":3333"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
