package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"87.GO/internal/models"
	"87.GO/internal/postgres"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	ItemService *postgres.Item
}

func NewItemHandler(db *sql.DB) *ItemHandler {
	return &ItemHandler{ItemService: postgres.NewItem(db)}
}

// CreateItem godoc
// @Summary Create a new item
// @Description Create a new item with the input payload
// @Tags items
// @Accept  json
// @Produce  json
// @Param item body models.Item true "Item to create"
// @Success 200 {object} models.Item
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items [post]
func (i *ItemHandler) CreateItem(c *gin.Context) {
	var item *models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	item, err := i.ItemService.StoreNewItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": item.ID})
}

// GetItemById godoc
// @Summary Get an item by ID
// @Description Get an item by ID
// @Tags items
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items/{id} [get]
func (i *ItemHandler) GetItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := i.ItemService.StoreGetItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// UpdateItem godoc
// @Summary Update an item
// @Description Update an item with the input payload
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param item body models.Item true "Item to update"
// @Success 200 {object} models.Item
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items/{id} [put]
func (i *ItemHandler) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var item *models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	item.ID = id

	item, err = i.ItemService.StoreUpdateItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// DeleteItem godoc
// @Summary Delete an item
// @Description Delete an item by ID
// @Tags items
// @Param id path int true "Item ID"
// @Success 204 {object} nil
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /items/{id} [delete]
func (i *ItemHandler)DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := i.ItemService.StoreDeleteItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
