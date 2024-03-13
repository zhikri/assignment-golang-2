package controllers

import (
	"assignment-two/models"
	"net/http"

	"assignment-two/database"
	"github.com/gin-gonic/gin"
)

// Create Order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

// Get Order
func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

// Update Order
func UpdateOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	var updatedOrder models.Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedOrder.ID = order.ID
	database.DB.Save(&updatedOrder)
	c.JSON(http.StatusOK, updatedOrder)
}

// Delete Order
func DeleteOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	database.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Success delete"})
}
