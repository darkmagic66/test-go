package main

import (
	m "echo/controller"
	_ "fmt"
	_ "log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from /")

	})

	e.GET("/getAllData", m.GetAllData)
	e.GET("/getData/:id", m.GetData)
	e.POST("/createData", m.CreateData)
	// log for what is it doing
	// Fatal : it will kill process
	e.Logger.Fatal(e.Start("localhost:3000"))
}
