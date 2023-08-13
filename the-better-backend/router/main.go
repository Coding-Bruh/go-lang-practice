package router

import (
    "../handlers"

    "github.com/gofiberfiber/v2"
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