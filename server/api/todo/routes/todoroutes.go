package todoroutes

import (
	"github.com/veloa/goweb/server/api/todo/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.Get("/api/todos", todocontroller.GetAll)
	e.Post("/api/todos", todocontroller.NewTodo)
	e.Delete("/api/todos/:id", todocontroller.RemoveTodo)
}
