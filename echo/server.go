package main

import (
	_ "fmt"
	_ "log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Test struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Test string `json:"test"`
	Year int64  `josn:"year"`
}

var arr_test = []Test{
	{
		ID:   1,
		Name: "John",
		Test: "dog",
		Year: 2002,
	},
	{
		ID:   2,
		Name: "Micheal",
		Test: "cat",
		Year: 2008,
	},
}

func getAllData(c echo.Context) error {
	y := c.QueryParam("year")

	if y == "" {
		return c.JSON(http.StatusOK, arr_test)
	}

	year, err := strconv.ParseInt(y, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	target := []Test{}
	for _, val := range arr_test {
		if val.Year == year {
			target = append(target, val)
		}
	}

	// header application is set to JSON
	return c.JSON(http.StatusOK, arr_test)

}

func getData(c echo.Context) error {
	id, _ := strconv.ParseInt((c.Param("id")), 0, 64)
	for _, val := range arr_test {
		if val.ID == id {
			return c.JSON(http.StatusOK, val)
		}
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "not found"})
}

func createData(c echo.Context) error {
	obj := new(Test)

	// ioutill.ReadAll  & json.Unmarshal (json -> struct)
	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	arr_test = append(arr_test, *obj)
	return c.JSON(http.StatusCreated, "create success")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from /")

	})

	e.GET("/getAllData", getAllData)
	e.GET("/getData/:id", getData)
	e.POST("/createData", createData)
	// log for what is it doing
	// Fatal : it will kill process
	e.Logger.Fatal(e.Start("localhost:3000"))
}
