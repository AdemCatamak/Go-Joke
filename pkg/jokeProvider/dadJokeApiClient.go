package jokeProvider

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type dadJokeApiClient struct {
	settings *dadJokeApiSettings
	client   *http.Client
}

func newDadJokeApiClient() *dadJokeApiClient {
	settings := newDadJokeApiSettings()

	client := &http.Client{}
	client.Timeout = time.Duration(settings.Timeout) * time.Millisecond

	return &dadJokeApiClient{
		settings: settings,
		client:   client,
	}
}

func (j dadJokeApiClient) GetJoke() string {
	req, err := http.NewRequest("GET", j.settings.BaseUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	res, err := j.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	log.Println("Response Body:")
	log.Println(string(body))

	if res.StatusCode != 200 {
		panic("Unexpected status code: " + res.Status)
	}

	var responseItems getJokeResponseItem
	err = json.Unmarshal(body, &responseItems)
	if err != nil {
		panic(err)
	}

	return responseItems.Joke
}

type getJokeResponseItem struct {
	Joke string `json:"joke"`
}
