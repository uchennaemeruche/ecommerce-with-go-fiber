package route

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/database"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/model"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func ProductResponse(productModel model.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func findProduct(id int, product *model.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("No product found with the given ID")
	}
	return nil
}

// Product godoc
// @Description create new product
// @Produce json
// @id CreateProduct
// @tag.name "Product"
// @tag.description "Product Routes"
// @Param Product formData Product true "Product input"
// @Success 201 {object} Product
// @Failure 400 {object}  HTTPError
// @Router /api/products/ [post]
func CreateProduct(c *fiber.Ctx) error {
	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	response := ProductResponse(product)
	return c.Status(200).JSON(response)
}

// Product godoc
// @Description Get all Products
// @Produce json
// @id GetProducts
// @tag.name "Product"
//  @tag.description "Product Routes"
// @Success 200 {object} []Product
// @Failure 400 {object}  HTTPError
// @Router /api/products/ [get]
func GetProducts(c *fiber.Ctx) error {
	products := []model.Product{}

	database.Database.Db.Find(&products)
	response := []Product{}

	for _, product := range products {
		response = append(response, ProductResponse(product))
	}

	return c.Status(200).JSON(response)
}

// Product godoc
// @Description Get a single Product
// @Produce json
// @Param id path int true "Product ID"
// @id GetProduct
// @tag.description "Product Routes"
// @tag.name "Product"
// @Success 200 {object} Product
// @Failure 400 {object}  HTTPError
// @Failure 404  {object}  HTTPError
// @Failure 500  {object}  HTTPError
// @Router /api/products/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Invalid product id")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(ProductResponse(product))
}

// Product godoc
// @Description Update a Product
// @Produce json
// @Param id path int true "Product ID"
// @Param Product formData Product true "Product input"
// @id UpdateProduct
// @tag.name "Product"
// @tag.description "Product Routes"
// @Success 200 {object} Product
// @Failure 400 {object}  HTTPError
// @Failure 404 {object}  HTTPError
// @Failure 500 {object}  HTTPError
// @Router /api/products/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Invalid Product ID")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	var updateData updateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	return c.Status(200).JSON(ProductResponse(product))
}

// Product godoc
// @Description Delete a Product
// @Produce json
// @Param id path int true "Product ID"
// @id DeleteProduct
// @tag.name "Product"
// @tag.description "Product Routes"
// @Success 200 {object} Product
// @Failure 400 {object}  HTTPError
// @Failure 404 {object}  HTTPError
// @Failure 500 {object}  HTTPError
// @Router /api/products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Invalid product ID")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(202).SendString("Product deleted")
}
