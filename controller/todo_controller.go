package controller

import (
	"github.com/gofiber/fiber/v2"

	"four-layer-todo-app/model"
	"four-layer-todo-app/service"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	todo := new(model.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	todoSaved, err := c.todoService.CreateTodo(ctx.Context(), todo)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(todoSaved)
}

func (c *TodoController) GetTodos(ctx *fiber.Ctx) error {
	todos, err := c.todoService.GetAllTodos(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch todos",
		})
	}

	return ctx.JSON(todos)
}
