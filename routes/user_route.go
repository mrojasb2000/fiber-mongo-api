package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrojasb2000/fiber-mongo-api/controllers"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Get("/user", controllers.GetAllUsers)

	app.Get("/user/:userId", controllers.GetAUser)

	app.Post("/user", controllers.CreateUser)

	app.Put("/user/:userId", controllers.EditAUser)

	app.Delete("/user/:userId", controllers.DeleteAUser)

}
