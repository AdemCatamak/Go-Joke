package handlers

import (
	"Go-Joke/pkg/jokeProvider"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetJoke(c *gin.Context) {
	jokeClient := jokeProvider.NewJokeApiClient()
	joke := jokeClient.GetJoke()

	c.JSON(http.StatusOK, gin.H{
		"joke": joke,
	})
}
