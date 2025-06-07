package routes

import (
	"github.com/arie0703/go-headless-cms-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// CORS設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.GET("/content-types", func(c *gin.Context) { controllers.GetContentTypes(c, db) })
		api.POST("/content-types", func(c *gin.Context) { controllers.CreateContentType(c, db) })
		api.GET("/users", func(c *gin.Context) { controllers.GetAllUsers(c, db) })
		api.POST("/users", func(c *gin.Context) { controllers.CreateUser(c, db) })
		api.PUT("/users/:id", func(c *gin.Context) { controllers.UpdateUser(c, db) })
		api.DELETE("/users/:id", func(c *gin.Context) { controllers.DeleteUser(c, db) })

		// Field routes
		api.GET("/fields", func(c *gin.Context) { controllers.GetFields(c, db) })
		api.POST("/fields", func(c *gin.Context) { controllers.CreateField(c, db) })
		api.PUT("/fields/:id", func(c *gin.Context) { controllers.UpdateField(c, db) })
		api.DELETE("/fields/:id", func(c *gin.Context) { controllers.DeleteField(c, db) })
	}

	return r
}
