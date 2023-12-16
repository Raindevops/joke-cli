package jokes

import (
	"encoding/json"
	"fmt"
	"joke-cli/internal/api"
	"log"
	"time"
)

const api_url = "https://official-joke-api.appspot.com/"

type Joke struct {
	Type  string `json:"type"`
	Setup string `json:"setup"`
	Joke  string `json:"punchline"`
	ID    int    `json:"id"`
}

func GetJoke() {
	rspBytes := api.GetDataFromApi(api_url + "random_joke")
	joke := Joke{}

	if err := json.Unmarshal(rspBytes, &joke); err != nil {
		log.Printf("Could not unmarchall responseBytes. %v", err)
	}

	fmt.Printf("%v \n", joke.Setup)
	time.Sleep(5 * time.Second)
	fmt.Printf(joke.Joke)
}
