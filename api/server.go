package api

import (
	"Go-Joke/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Recovery())

	registerRoutes(engine)

	return engine
}
