package cmd

import (
	"log"
	"todo-go/config"
	"todo-go/database"
	"todo-go/pkg/logger"
	"todo-go/routes"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	app.Use(logger.LoggerMiddleware())

	config.LoadEnv()

	database.ConnectDB()

	routes.RegisterRoutes(app)

	port := config.GetEnv("PORT", "4000")

	log.Fatal(app.Listen(":" + port))
}
