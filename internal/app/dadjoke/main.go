package dadjoke

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const api_url = "https://icanhazdadjoke.com/"

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type SearchResult struct {
	Results    json.RawMessage `json:"results"`
	SearchTerm string          `json:"search_term"`
	Status     int             `json:"status"`
	TotalJokes int             `json:"total_jokes"`
}

func GetRandomJoke() {
	responseBytes := getJokeData(api_url)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Could not unmarshal responseByte. %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	r, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("Cloud not request a dadjoke. %v", err)
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("User-agent", "Joke CLI (https://github.com/Raindevops/joke-cli)")

	rsp, err := http.DefaultClient.Do(r)

	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := io.ReadAll(rsp.Body)

	if err != nil {
		log.Printf("Could not read response Body. %v", err)
	}

	return responseBytes
}

func GetRandomJokeWithTerm(jokeTerm string) {
	total, results := getJokeDataWithTerm(jokeTerm)
	randomiseJokeList(total, results)
}

func getJokeDataWithTerm(jokeTerm string) (totalJokes int, jokeList []Joke) {
	responseBytes := getJokeData(api_url + "/search?term=" + jokeTerm)

	jokeListRaw := SearchResult{}

	if err := json.Unmarshal(responseBytes, &jokeListRaw); err != nil {
		log.Printf("Could not unmarshal response. %v", err)
	}
	jokes := []Joke{}
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Printf("Could not unmarshal response. %v", err)
	}

	return jokeListRaw.TotalJokes, jokes
}

func randomiseJokeList(length int, jokeList []Joke) {
	rand.NewSource(time.Now().Unix())

	min := 0
	max := length - 1

	if length <= 0 {
		err := fmt.Errorf("no jokes found with this term")
		fmt.Println(err.Error())
	} else {
		randomNum := min + rand.Intn(max-min)
		fmt.Println(jokeList[randomNum].Joke)
	}
}
