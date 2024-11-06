package jokeProvider

type JokeApiClient interface {
	GetJoke() string
}

func NewJokeApiClient() JokeApiClient {
	client := newDadJokeApiClient()
	return client
}
