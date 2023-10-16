package routers

import (
	"easy-rbac/handlers"
	"easy-rbac/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware(db))
	r.Use(middleware.CheckPermission(db))

	api := r.Group("/api")
	{
		api.GET("/v1/user", handlers.GetUser) // just a sample route
	}

	return r
}
