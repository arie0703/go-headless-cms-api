package controllers

import (
	"net/http"

	"github.com/arie0703/go-headless-cms-api/models"
	"github.com/arie0703/go-headless-cms-api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetContentTypes(c *gin.Context, db *gorm.DB) {
	var types []models.ContentType
	if err := db.Find(&types).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "取得失敗"})
		return
	}
	c.JSON(http.StatusOK, types)
}

func CreateContentType(c *gin.Context, db *gorm.DB) {
	var ct models.ContentType
	if err := c.ShouldBindJSON(&ct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不正なデータ"})
		return
	}
	ct.ID = utils.GenerateUUID()
	if err := db.Create(&ct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "作成失敗"})
		return
	}
	c.JSON(http.StatusCreated, ct)
}
