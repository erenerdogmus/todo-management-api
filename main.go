package main

import (
	"github.com/erenerdogmus/app"
	"github.com/erenerdogmus/configs"
	"github.com/erenerdogmus/repository"
	"github.com/erenerdogmus/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.ConnectDB(), "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDb(dbClient)
	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":8080")
}
