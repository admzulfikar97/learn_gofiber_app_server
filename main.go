package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) {
	// 	c.SendString("API GoFiber UP!!!")
	// })
	var err error
	app.Get("/", func(c *fiber.Ctx) error {
		err = c.SendString("API GoFiber UP!!!")
		return err
	})
	app.Get("/home", func(c *fiber.Ctx) error {
		err = c.SendString("--API HOME SERVER--")
		return err
	})

	app.Listen(":9077")
}
