package api

import (
	"Go-Joke/api/handlers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(api *gin.RouterGroup) {
	api.GET("", handlers.GetIndex)
	api.GET("index", handlers.GetIndex)

	api.GET("joke", handlers.GetJoke)
}
