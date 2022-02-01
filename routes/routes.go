package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vargaae/go-auth/controllers"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)

}
