package handlers

import (
	"context"
	"net/http"
	"strconv"

	"ApiRest/model"

	"github.com/labstack/echo/v4"
)

type TaskUsecase interface {
	GetTaskByID(ctx context.Context, id uint) ([]model.Task, error)
	GetAllTask(ctx context.Context) ([]model.Task, error)
	UpdateTask(ctx context.Context, id uint, p *model.Task) error
	DeleteTask(ctx context.Context, id uint) error
	AddTask(ctx context.Context, task *model.Task) error
}

type TaskHandler struct {
	TUsecase TaskUsecase
}

func NewTaskHandler(e *echo.Echo, us TaskUsecase) {
	handler := &TaskHandler{
		TUsecase: us,
	}
	e.GET("/tasks/:id", handler.GetByID)
	e.GET("/tasks/", handler.GetAllTask)
	e.POST("/tasks/", handler.AddTask)
	e.PUT("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)
}

func (t *TaskHandler) GetAllTask(c echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := t.TUsecase.GetAllTask(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка (500 Internal Server Error): Проблема на сервере.")
	}

	return c.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Ошибка (404 Not Found): Задача не найдена.")
	}

	ctx := c.Request().Context()

	task, err := t.TUsecase.GetTaskByID(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка (500 Internal Server Error): Проблема на сервере.")
	}

	return c.JSON(http.StatusOK, task)
}

func (t *TaskHandler) AddTask(c echo.Context) (err error) {
	var task model.Task
	err = c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Ошибка (400 Bad Request): Неверный формат данных")
	}

	ctx := c.Request().Context()
	err = t.TUsecase.AddTask(ctx, &task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка (500 Internal Server Error): Не удалось обновить задачу")
	}

	return c.JSON(http.StatusCreated, task)
}

func (t *TaskHandler) UpdateTask(c echo.Context) (err error) {
	var input model.Task
	if err1 := c.Bind(&input); err1 != nil {
		return c.JSON(http.StatusBadRequest, "Ошибка (400 Bad Request): Неверный формат данных")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Ошибка (404 Not Found): Задача не найдена")
	}

	ctx := c.Request().Context()
	if err2 := t.TUsecase.UpdateTask(ctx, uint(id), &input); err2 != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка (500 Internal Server Error): Не удалось обновить задачу")
	}

	return c.JSON(http.StatusOK, "Успех (200 OK):")
}

func (t *TaskHandler) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Ошибка (404 Not Found): Задача не найдена.")
	}

	ctx := c.Request().Context()
	err = t.TUsecase.DeleteTask(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка (500 Internal Server Error): Проблема на сервере")
	}

	return c.JSON(http.StatusNoContent, "Успех (204 No Content): Задача удалена.")
}
