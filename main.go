package main

import (
	"fmt"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	cfg := basicauth.Config{
		Users: map[string]string{
			"admin": "password123",
		},
	}
	app.Use(basicauth.New(cfg))

	app.Get("/:name?", func(c *fiber.Ctx) {
		if c.Params("name") != "" {
			c.Send("Hello " + c.Params("name") + "!")
		}
	})

	fmt.Println("Application will be running on http://localhost")
	app.Listen(80)
}
