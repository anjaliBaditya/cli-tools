package main

import (
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// User represents a user object
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users  = make(map[int]User)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	app := fiber.New()

	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	app.Listen(":8080")
}

func getUsers(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	return c.JSON(userList)
}

func getUser(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	user.ID = nextID
	nextID++
	users[user.ID] = user

	return c.Status(fiber.StatusCreated).JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	_, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	updatedUser.ID = id
	users[id] = updatedUser

	return c.JSON(updatedUser)
}

func deleteUser(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	_, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	delete(users, id)
	return c.SendStatus(fiber.StatusNoContent)
}
