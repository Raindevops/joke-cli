package jokes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const api_url = "https://official-joke-api.appspot.com/"

type Joke struct {
	Type  string `json:"type"`
	Setup string `json:"setup"`
	Joke  string `json:"punchline"`
	ID    int    `json:"id"`
}

func getJokesData(uri string) []byte {
	r, err := http.NewRequest(
		http.MethodGet,
		uri,
		nil,
	)

	if err != nil {
		log.Printf("Could not create a new request to the api. %v", err)
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("User-agent", "Joke CLI (https://github.com/Raindevops/Joke-cli)")

	rsp, err := http.DefaultClient.Do(r)

	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	rspBytes, err := io.ReadAll(rsp.Body)

	if err != nil {
		log.Printf("Could not read response Body. %v", err)
	}

	return rspBytes
}

func GetJoke() {
	rspBytes := getJokesData(api_url + "random_joke")
	joke := Joke{}

	if err := json.Unmarshal(rspBytes, &joke); err != nil {
		log.Printf("Could not unmarchall responseBytes. %v", err)
	}

	fmt.Printf("%v \n", joke.Setup)
	time.Sleep(5 * time.Second)
	fmt.Printf(joke.Joke)
}
