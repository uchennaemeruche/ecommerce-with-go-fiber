package route

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/database"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/model"
)

// Some sor of serializer
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func CreateResponseUser(userModel model.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, Lastname: userModel.Lastname}
}

func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(201).JSON(responseUser)

}

func GetUsers(c *fiber.Ctx) error {
	users := []model.User{}

	database.Database.Db.Find(&users)
	resoponseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		resoponseUsers = append(resoponseUsers, responseUser)
	}
	return c.Status(200).JSON(resoponseUsers)
}

func findUser(id int, user *model.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	var user model.User

	if err != nil {
		return c.Status(400).JSON("Please ensure you pass a valid user id")
	}

	if err := findUser(userId, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	var user model.User

	if err != nil {
		return c.Status(400).JSON("Please ensure you pass a valid user id")
	}

	if err := findUser(userId, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updateUser struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastName"`
	}

	var updateData updateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.Lastname = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	var user model.User

	if err != nil {
		return c.Status(400).JSON("Please ensure you pass a valid user id")
	}

	if err := findUser(userId, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("User deleted")

}
