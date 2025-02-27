package main

import (
	"jokenpo/internal/handlers"
	"jokenpo/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/jogar", middleware.AuthMiddleware(), handlers.PlayJokenpo)

	r.Run(":8080")
}
