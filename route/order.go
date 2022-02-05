package route

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/database"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/model"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func OrderResponse(order model.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order model.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user model.User
	if err := findUser(order.UserID, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product model.Product
	if err := findProduct(order.ProductID, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	return c.Status(201).JSON(OrderResponse(order, CreateResponseUser(user), ProductResponse(product)))

}

func GetOrders(c *fiber.Ctx) error {
	orders := []model.Order{}
	// var user model.User
	// var product model.Product

	database.Database.Db.Find(&orders)

	response := []Order{}

	for _, order := range orders {
		fmt.Println(order)
		var user model.User
		database.Database.Db.Find(&user, "id = ?", order.UserID)
		var product model.Product
		database.Database.Db.Find(&product, "id = ?", order.ProductID)
		response = append(response, OrderResponse(order, CreateResponseUser(user), ProductResponse(product)))
	}

	return c.Status(200).JSON(response)
}
