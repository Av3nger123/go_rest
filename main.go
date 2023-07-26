package main

import (
	"todo-app/config"
	"todo-app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {

	app := fiber.New()

	// Initialize MongoDB connection
	err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer config.Disconnect()

	// middlewares
	app.Use(cors.New())
	app.Use(cache.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Get("/todo", handlers.GetAllTodo)
	app.Get("/todo/:id", handlers.GetTodoById)
	app.Post("/todo", handlers.CreateTodo)
	app.Put("/todo/:id", handlers.UpdateTodo)
	app.Delete("/todo/:id", handlers.DeleteTodo)

	// Start the server
	app.Listen(":3000")

}
