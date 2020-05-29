package main

import (
	"jwt-demo/internal/middleware"
	"jwt-demo/internal/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	if err != nil {
		log.Fatalf("auth middleware error: %s", err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth := r.Group("/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/userinfo", user.Info)
	}

	r.Run()
}
