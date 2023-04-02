package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/thegeorgenikhil/go-for-nodejs-devs/internal/handlers"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Hello from middleware")
		return c.Next()
	})

	app.Get("/healthcheck", healthcheck)

	app.Get("/api/products", handlers.GetAllProducts)
	app.Post("/api/products", handlers.CreateProduct)

	log.Fatal(app.Listen(":3000"))
}
