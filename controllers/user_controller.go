package controllers

import (
	"net/http"

	"github.com/arie0703/go-headless-cms-api/models"
	"github.com/arie0703/go-headless-cms-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザー一覧取得に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = utils.GenerateUUID() // UUIDを自動生成

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB作成失敗"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var user models.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "削除失敗"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "削除完了"})
}
