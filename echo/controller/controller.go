package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// -------------------------------------------------------
type Test struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Test string `json:"test"`
	Year int64  `josn:"year"`
}

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Owl struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// -------------------------------------------------------

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

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]

func HeatCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "statsu ok",
	})
}

func GetAllData(c echo.Context) error {
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

func GetData(c echo.Context) error {
	id, _ := strconv.ParseInt((c.Param("id")), 0, 64)
	for _, val := range arr_test {
		if val.ID == id {
			return c.JSON(http.StatusOK, val)
		}
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "not found"})
}

func CreateData(c echo.Context) error {
	obj := new(Test)

	// echo.bind -> ioutill.ReadAll  & json.Unmarshal (json -> struct)
	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	arr_test = append(arr_test, *obj)
	return c.JSON(http.StatusCreated, "create success")
}

//-----------------------------------------------------------------------------------------

// pure but fastest
func AddCat(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()

	// read body
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading request  %s\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s\n", err))
	}

	// json -> obj
	err = json.Unmarshal(body, &cat)
	if err != nil {
		log.Printf("Failed unmashal %s\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s\n", err))
	}

	// success
	log.Printf("This is your cat %v", cat)
	return c.String(http.StatusOK, "We got your cat")
}

func AddDog(c echo.Context) error {
	dog := Dog{}
	defer c.Request().Body.Close()

	// read by json decode to obj assign to dog
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Fail to convert to json and decote to obj %s\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s\n", err))
	}

	// success
	log.Printf("This is your dog %v", dog)
	return c.String(http.StatusOK, "We got your dog")
}

func AddOwl(c echo.Context) error {
	owl := Owl{}
	defer c.Request().Body.Close()

	// read by json decode to obj assign to dog
	err := c.Bind(&owl)
	if err != nil {
		log.Printf("Fail to bind %s\n", err)
	}

	// success
	log.Printf("This is your owl %v", owl)
	return c.String(http.StatusOK, "We got your owl")
}

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the admin page")
}
