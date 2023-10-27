package main

import (
	"github.com/admzulfikar97/learn_gofiber_app_server/database"
	"github.com/admzulfikar97/learn_gofiber_app_server/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to DB
	database.ConnectDB()

	routes.SetupRoutes(app)

	app.Listen(":9077")
}
