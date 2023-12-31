package app

import (
    "os"

    "github.com/Coding-Bruh/go-lang-practice/tree/main/the-better-backend/config"
    "github.com/Coding-Bruh/go-lang-practice/tree/main/the-better-backend/database"
    "github.com/Coding-Bruh/go-lang-practice/tree/main/the-better-backend/router"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() error {
    // load env
    err := config.LoadENV()
    if err != nil {
        return err
    }

    // start database
    err = database.StartMongoDB()
    if err != nil {
        return err
    }

    // defer closing database
    defer database.CloseMongoDB()

    // create app
    app := fiber.New()

    // attach middleware
    app.Use(recover.New())
    app.Use(logger.New(logger.Config{
        format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
    }))

    // setup routes
    router.SetupRoutes(app)

    // attach swagger
    config.AddSwaggerRoutes(app)

    // get the port and start
    port := os.Getenv("PORT")
    app.Listem(":" + port)

    return nil
}
