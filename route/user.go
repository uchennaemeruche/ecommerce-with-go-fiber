package route

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/database"
	"github.com/uchennaemeruche/ecommerce-with-go-fiber/model"
)

// Some sort of serializer
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type HTTPError struct {
	status  string
	message string
}

func CreateResponseUser(userModel model.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, Lastname: userModel.Lastname}
}

// User godoc
// @Description create new user
// @Produce json
// @id CreateUser
// @tag.name "User"
// @Param User formData User true "user input"
// @Success 201 {object} User
// @Failure 400 {object}  HTTPError
// @Router /api/users/ [post]
func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(201).JSON(responseUser)

}

// User godoc
// @Description Get all users
// @Produce json
// @id GetUsers
// @tag.name User
// @Success 200 {object} []User
// @Failure 400 {object}  HTTPError
// @Router /api/users/ [get]
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

// User godoc
// @Description Get a single user
// @Produce json
// @Param id path int true "User ID"
// @id GetUser
// @tag.name User
// @Success 200 {object} User
// @Failure 400 {object}  HTTPError
// @Failure 404  {object}  HTTPError
// @Failure 500  {object}  HTTPError
// @Router /api/users/{id} [get]
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

// User godoc
// @Description Update a user
// @Produce json
// @Param id path int true "User ID"
// @Param User formData User true "user input"
// @id UpdateUser
// @tag.name User
// @Success 200 {object} User
// @Failure 400 {object}  HTTPError
// @Failure 404 {object}  HTTPError
// @Failure 500 {object}  HTTPError
// @Router /api/users/{id} [put]
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

// User godoc
// @Description Delete a user
// @Produce json
// @Param id path int true "User ID"
// @id DeleteUser
// @tag.name User
// @Success 200 {object} User
// @Failure 400 {object}  HTTPError
// @Failure 404 {object}  HTTPError
// @Failure 500 {object}  HTTPError
// @Router /api/users/{id} [delete]
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
