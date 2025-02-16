package handlers

import (
	"todo-go/internals/models"
	"todo-go/internals/services"
	"todo-go/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	TodoService *services.TodoService
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{TodoService: services.NewTodoService()}
}

func (h *TodoHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON(response.ErrorResponse("the id is not int"))
	}

	todo, err := h.TodoService.Get(id)
	if err != nil {
		return c.Status(200).JSON(response.ErrorResponse("not to with this id"))
	}

	return c.Status(200).JSON(response.SuccessResponse("one user", todo))
}

func (h *TodoHandler) GetAll(c *fiber.Ctx) error {
	todos, err := h.TodoService.GetAllTodos()

	if err != nil {
		return c.Status(200).JSON(response.ErrorResponse("no todos"))
	}
	return c.Status(200).JSON(response.SuccessResponse("list of all user", todos))
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	todo := models.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("no data found"))
	}

	if todo.Title == "" {
		return c.Status(404).JSON(fiber.Map{"error": "body is empty"})
	}

	err := h.TodoService.CreateTodo(&todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.SuccessResponse("todo created", todo))
	}

	return c.Status(200).JSON(response.SuccessResponse("to creates successfully ", todo))
}

func (h *TodoHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON(response.ErrorResponse("the id is not int"))
	}

	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(200).JSON(response.ErrorResponse("the id is not int"))
	}
	todo.ID = uint(id)
	err = h.TodoService.UpdateTodo(&todo)
	if err != nil {
		return c.Status(200).JSON(response.ErrorResponse("failed to update"))
	}

	return c.Status(200).JSON(response.SuccessResponse("to creates successfully ", todo))
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(200).JSON(response.ErrorResponse("the id is not int"))
	}

	err = h.TodoService.Delete(id)
	if err != nil {
		return c.Status(200).JSON(response.ErrorResponse("failed to delete"))
	}

	return c.Status(200).JSON(response.SuccessResponse("deleted successfully ", nil))

}
