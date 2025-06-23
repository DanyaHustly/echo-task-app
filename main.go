package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var task string

type RequestBody struct {
	Task string `json:"task"`
}

func postTaskHandler(c echo.Context) error {
	var reqBody RequestBody

	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	task = reqBody.Task

	return c.String(http.StatusOK, "Задача получена")
}

func getTaskHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Задача: "+task)
}
func main() {
	e := echo.New()
	e.POST("/task", postTaskHandler)
	e.GET("/", getTaskHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
