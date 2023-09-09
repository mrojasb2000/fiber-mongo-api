package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrojasb2000/fiber-mongo-api/configs"
	"github.com/mrojasb2000/fiber-mongo-api/routes"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app) //add this

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "Welcome to Fiber Framework",
		})
	})

	app.Listen(":3000")
}
