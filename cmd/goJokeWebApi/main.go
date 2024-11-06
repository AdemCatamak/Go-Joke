package main

import (
	"Go-Joke/api"
	"Go-Joke/internal/config"
	"fmt"
	"log"
)

func main() {
	log.Println("App Is Started")

	configManager := config.GetConfigManager()
	portNumber := configManager.GetString(config.PortKey)

	s := api.NewServer()

	err := s.Run(":" + portNumber)
	if err != nil {
		panic(fmt.Sprintf("Server cannot start: %s", err))
	}
}
