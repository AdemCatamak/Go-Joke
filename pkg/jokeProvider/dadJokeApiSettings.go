package jokeProvider

import "Go-Joke/internal/config"

type dadJokeApiSettings struct {
	BaseUrl string
	Timeout int
}

func newDadJokeApiSettings() *dadJokeApiSettings {
	configManager := config.GetConfigManager()
	settings := &dadJokeApiSettings{}
	configManager.GetObj("JokeApiSettings", settings)

	return settings
}
