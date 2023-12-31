package router

import (
    "github.com/Coding-Bruh/go-lang-practice/tree/main/the-better-backend/handlers"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/health", handlers.HandleHealthCheck)

    // setup the todos group
    todos := app.Group("/todos")
    todos.Get("/", handlers.HandleAllTodos)
    todos.Post("/", handlers.HandleCreateTodo)
    todos.Put("/:id", handlers.HandleUpdateTodo)
    todos.Get("/:id", handlers.HandleGetTodo)
    todos.Delete("/:id", handlers.HandleDeleteTodo)+
}