package chucknorris

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const api_url = "https://api.chucknorris.io/jokes/"

type Phrase struct {
	ID     string `json:"id"`
	Phrase string `json:"value"`
}

type category string

type Categories []category

func RandomPhrase() {
	responseBytes := getRandomPhrase()
	phrase := Phrase{}
	if err := json.Unmarshal(responseBytes, &phrase); err != nil {
		fmt.Printf("Could not unmarshal responseByte. %v", err)
	}

	fmt.Println(string(phrase.Phrase))
}

func getRandomPhrase() []byte {
	r, err := http.NewRequest(
		http.MethodGet,
		api_url+"random",
		nil,
	)

	if err != nil {
		log.Printf("Could not get a Chuck Norris phrase. %v", err)
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

func ListAllCategories() {

}

func GetPhraseByCategory() {

}

// radnom joke
// list categories
// joke by category
