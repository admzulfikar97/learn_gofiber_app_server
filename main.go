package main

import (
	"fmt"

	"github.com/admzulfikar97/learn_gofiber_app_server/config"
	"github.com/admzulfikar97/learn_gofiber_app_server/database"
	"github.com/admzulfikar97/learn_gofiber_app_server/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to DB
	database.ConnectDB()

	// Routing
	var err error
	app.Get("/", func(c *fiber.Ctx) error {
		err = c.SendString("")
		p := config.Config("DB_PORT")
		err = c.SendString(`--localhost:9077--
		--API GoFiber UP!!!--`)
		fmt.Println("PORT:", p)
		return err
	})
	// app.Get("/home", func(c *fiber.Ctx) error {
	// 	err = c.SendString("--API HOME SERVER--")
	// 	return err
	// })
	// app.Get("/health", func(c *fiber.Ctx) error {
	// 	err = c.SendString("--Server localhost:9077 UP--")
	// 	return err
	// })
	router.SetupRoutes(app)

	app.Listen(":9077")
}
