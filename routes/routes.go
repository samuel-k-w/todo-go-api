package routes

import (
	"todo-go/internals/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	todoHandler := handlers.NewTodoHandler()

	api.Get("/todos", todoHandler.GetAll)
	api.Get("/todos/:id", todoHandler.Get)
	api.Post("/todos", todoHandler.Create)
	api.Patch("/todos/:id", todoHandler.Update)
	api.Delete("/todos/:id", todoHandler.Delete)
}
