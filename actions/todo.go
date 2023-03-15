package actions

import (
	"net/http"
	"project1/models"

	"github.com/gobuffalo/buffalo"
)

// TodoIndex default implementation.
func TodoIndex(c buffalo.Context) error {

	// Create an array to receive todos
	todos := []models.Todo{}

	// Get all the todos from database
	err := models.DB.All(&todos)

	// Handle any error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}

	// Return list of todos as json
	return c.Render(http.StatusOK, r.JSON(todos))
}

// TodoShow default implementation.
func TodoShow(c buffalo.Context) error {

	// Grab the id url parameter defined in app.go
	id := c.Param("id")

	// Create a variable to receive the todo
	todo := models.Todo{}

	// Grab the todo from the database
	err := models.DB.Find(&todo, id)

	// Handle possible error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}

	// Return the data as json
	return c.Render(http.StatusOK, r.JSON(&todo))
}

// TodoAdd default implementation.
func TodoAdd(c buffalo.Context) error {

	// Get item from url query
	item := c.Param("item")

	// Create new instance of todo
	todo := models.Todo{Item: item}

	// Create a todo without running validations
	err := models.DB.Create(&todo)

	// Handle error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}

	// Return new todo as json
	return c.Render(http.StatusOK, r.JSON(todo))
}
