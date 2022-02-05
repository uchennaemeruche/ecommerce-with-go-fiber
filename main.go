package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/database"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/route"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the app")
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", welcome)
	userRoutes := api.Group("/users", func(c *fiber.Ctx) error {
		// middleware for /api/users routes
		c.Set("Version", "v1")
		return c.Next()
	})
	userRoutes.Get("/", route.GetUsers)
	userRoutes.Get("/:id", route.GetUser)
	userRoutes.Post("/", route.CreateUser)
	userRoutes.Put("/:id", route.UpdateUser)
	userRoutes.Delete("/:id", route.DeleteUser)

	productRoutes := api.Group("/products", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	productRoutes.Get("/", route.GetProducts)
	productRoutes.Get("/:id", route.GetProduct)
	productRoutes.Post("/", route.CreateProduct)
	productRoutes.Put("/:id", route.UpdateProduct)
	productRoutes.Delete("/:id", route.DeleteProduct)

	orderRoutes := api.Group("/orders", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	orderRoutes.Post("/", route.CreateOrder)
	orderRoutes.Get("/", route.GetOrders)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	// app.Get("/", welcome)

	log.Fatal(app.Listen(":4000"))
}
