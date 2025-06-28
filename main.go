package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

var task = map[string]string{}

type RequestBody struct {
	Task string `json:"task"`
}

func postTaskHandler(c echo.Context) error {
	var reqBody RequestBody

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Неверный JSON"})
	}

	id := uuid.New().String()

	task[id] = reqBody.Task

	return c.JSON(http.StatusOK, map[string]string{
		"id":   id,
		"task": reqBody.Task,
	})
}

func getAllTasksHandler(c echo.Context) error {
	// Просто возвращаем всю карту задач в виде JSON
	return c.JSON(http.StatusOK, tasks)
}

func patchTaskHandler(c echo.Context) error {
	id := c.Param("id")

	if _, ok := task[id]; ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
	}

	var reqBody RequestBody
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Не верный JSON"})
	}

	task[id] = reqBody.Task

	return c.JSON(http.StatusOK, map[string]string{
		"id":   id,
		"task": reqBody.Task,
	})
}

func deleteTaskHandler(c echo.Context) error {
	id := c.Param("id")

	if _, ok := task[id]; !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
	}

	delete(task, id)

	return c.JSON(http.StatusOK, "Задача удалена")
}
func main() {
	e := echo.New()

	// Создать задачу (Create)
	e.POST("/tasks", postTaskHandler)

	// Получить все задачи (Read)
	e.GET("/tasks", getAllTasksHandler)

	// Обновить задачу по ID (Update)
	e.PATCH("/tasks/:id", patchTaskHandler)

	// Удалить задачу по ID (Delete)
	e.DELETE("/tasks/:id", deleteTaskHandler)

	// Запускаем сервер на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}
