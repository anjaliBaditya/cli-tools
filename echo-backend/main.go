package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func getUsers(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	return c.JSON(http.StatusOK, userList)
}

func getUser(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	user, exists := users[id]
	if !exists {
		return c.String(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	var user User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	user.ID = nextID
	nextID++
	users[user.ID] = user

	return c.JSON(http.StatusCreated, user)
}

func updateUser(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	_, exists := users[id]
	if !exists {
		return c.String(http.StatusNotFound, "User not found")
	}

	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	updatedUser.ID = id
	users[id] = updatedUser

	return c.JSON(http.StatusOK, updatedUser)
}

func deleteUser(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	_, exists := users[id]
	if !exists {
		return c.String(http.StatusNotFound, "User not found")
	}

	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
