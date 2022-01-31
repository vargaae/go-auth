package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:rootroot@/smart_brain"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// routes.Setup(app)

	app.Listen(":8000")
}
