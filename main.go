package main

import (
	"github.com/gin-gonic/gin"

	Handler "github.com/Spr1zze/barber-shop-backend/internal/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/users", Handler.GetUsers)
	router.POST("/users", Handler.PostUsers)

	router.Run("localhost:8000")
}
