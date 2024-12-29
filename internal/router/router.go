package router

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/CloudNavX/internal/auth"
	redis_db "github.com/harsh082ip/CloudNavX/internal/db/redis"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize Redis
	redis_db.InitializeRedis()

	// Setup OAuth authentication
	auth.SetupGoogleAuth()

	// Register Auth Routes
	authGroup := router.Group("/auth")
	auth.RegisterRoutes(authGroup)

	return router
}
