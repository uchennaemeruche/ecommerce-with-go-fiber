package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/database"
	_ "github.com/uchennaemeruche/ecommerce-with-go-fiber/docs"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/route"
)

// Welcome godoc
// @Description show the welcome page
// @Success 200 {string} string
// @Failure 400 {object}  HTTPError
// @Router /api/ [get]
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
	orderRoutes.Get("/", route.GetOrders2)
}

// @title Ecommerce Api using Go Fiber
// @version 1.0
// @description Go-Fiber implementation of the e-commerce
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email mail.asktech@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	setupRoutes(app)

	// app.Get("/", welcome)

	log.Fatal(app.Listen(":4000"))
}

type HTTPError struct {
	status  string
	message string
}
