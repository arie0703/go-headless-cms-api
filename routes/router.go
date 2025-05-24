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
	}

	return r
}
