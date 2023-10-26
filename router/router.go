package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// app := fiber.App()
	api := app.Group("/api")

	user := api.Group("/user")
	user.Get("/userId", func(c *fiber.Ctx) error {
		err := c.SendString("API api/user/userId")
		return err
	})

	note := api.Group("/note")
	note.Get("/noteId", func(c *fiber.Ctx) error {
		err := c.SendString("API api/note/noteId")
		return err
	})
}
