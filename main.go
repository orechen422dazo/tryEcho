package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Item struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Purchased bool    `json:"purchased"`
}

var items []Item
var lastID = 0

func main() {
	e := echo.New()

	// Routes
	e.GET("/items", getItems)
	e.GET("/items/:id", getItem)
	e.POST("/items", createItem)
	e.PUT("/items/:id", updateItem)
	e.DELETE("/items/:id", deleteItem)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// getItems returns all items
func getItems(c echo.Context) error {
	return c.JSON(http.StatusOK, items)
}

// getItem returns a specific item by id
func getItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	for _, item := range items {
		if item.ID == id {
			return c.JSON(http.StatusOK, item)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
}

// createItem adds a new item to the list
func createItem(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	lastID++
	item.ID = lastID
	items = append(items, *item)

	return c.JSON(http.StatusCreated, item)
}

// updateItem updates an existing item
func updateItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	item := new(Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	for i, existingItem := range items {
		if existingItem.ID == id {
			item.ID = id
			items[i] = *item
			return c.JSON(http.StatusOK, item)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
}

// deleteItem removes an item from the list
func deleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Item deleted successfully"})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
}
