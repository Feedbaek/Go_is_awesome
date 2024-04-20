package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//fmt.Println("Server is running on port 1323")
	//e.Logger.Fatal(e.Start(":1323"))
	go e.Start(":1323")
	go e.Start(":1324")
	e.Start(":1325")
}
