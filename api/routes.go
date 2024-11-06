package api

import (
	"Go-Joke/api/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(engine *gin.Engine) {
	api := engine.Group("api")

	api.GET("", handlers.GetIndex)
	api.GET("index", handlers.GetIndex)

	api.GET("joke", handlers.GetJoke)

	engine.LoadHTMLFiles("statics/index.html")
	serveIndexHtml := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
	engine.GET("/", serveIndexHtml)
	engine.GET("/index", serveIndexHtml)
}
