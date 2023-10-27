package routes

import (
	"fmt"

	"github.com/admzulfikar97/learn_gofiber_app_server/config"
	noteHandler "github.com/admzulfikar97/learn_gofiber_app_server/internal/handlers/note"
	"github.com/admzulfikar97/learn_gofiber_app_server/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
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

	// app := fiber.App()
	api := app.Group("/api", logger.New())

	user := api.Group("/user")
	user.Get("/userId", func(c *fiber.Ctx) error {
		err := c.SendString("API api/user/userId")
		return err
	})

	// API + Embedded function
	// Get parameter from url
	user.Get("/:param", func(c *fiber.Ctx) error {
		err := c.SendString("API : api/user/:param")
		param := c.Params("param")
		fmt.Println("parameter", param)
		allParams := c.AllParams()
		fmt.Println("All Params", allParams)
		return err
	})

	// Get parameter from url + get body request
	user.Post("/", func(c *fiber.Ctx) error {
		err := c.SendString("API : api/user")
		request := c.Request()
		fmt.Println("request", request)
		userRequestBody := new(model.User)
		c.BodyParser(userRequestBody)
		fmt.Println(userRequestBody.ID, userRequestBody.UserName, userRequestBody.Gender)
		err = c.Send(c.Body())
		return err
	})

	note := api.Group("/note")
	note.Get("/noteId", func(c *fiber.Ctx) error {
		err := c.SendString("API api/note/noteId")
		return err
	})

	// API
	note.Get("/", noteHandler.GetNotes)
	note.Post("/", noteHandler.CreateNotes)
	note.Get("/:noteID", noteHandler.GetNote)
	note.Put("/:noteID", noteHandler.UpdateNote)
	note.Delete("/:noteID", noteHandler.DeleteNote)
}
