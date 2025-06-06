package routes

import (
	"github.com/arie0703/go-headless-cms-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/content-types", func(c *gin.Context) { controllers.GetContentTypes(c, db) })
		api.POST("/content-types", func(c *gin.Context) { controllers.CreateContentType(c, db) })
		api.GET("/users", func(c *gin.Context) { controllers.GetAllUsers(c, db) })
		api.POST("/users", func(c *gin.Context) { controllers.CreateUser(c, db) })
		api.PUT("/users/:id", func(c *gin.Context) { controllers.UpdateUser(c, db) })
		api.DELETE("/users/:id", func(c *gin.Context) { controllers.DeleteUser(c, db) })
	}

	return r
}
