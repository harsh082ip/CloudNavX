package auth

import "github.com/gin-gonic/gin"

// RegisterRoutes registers the authentication routes
func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/google", LoginHandler)
	router.GET("/google/callback", CallbackHandler)
}
