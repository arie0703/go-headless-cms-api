package controllers

import (
	"net/http"

	"github.com/arie0703/go-headless-cms-api/models"
	"github.com/arie0703/go-headless-cms-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFields(c *gin.Context, db *gorm.DB) {
	var fields []models.Field
	if err := db.Find(&fields).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fields)
}

func CreateField(c *gin.Context, db *gorm.DB) {
	var field models.Field
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	field.ID = utils.GenerateUUID()
	if err := db.Create(&field).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, field)
}

func UpdateField(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var field models.Field

	if err := db.First(&field, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}

	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&field).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, field)
}

func DeleteField(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var field models.Field

	if err := db.First(&field, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}

	if err := db.Delete(&field).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Field deleted successfully"})
}
