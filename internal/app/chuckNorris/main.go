package chucknorris

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const api_url = "https://api.chucknorris.io/jokes/"

type Phrase struct {
	ID     string `json:"id"`
	Phrase string `json:"value"`
}

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
	r, err := http.NewRequest(
		http.MethodGet,
		api_url+"categories",
		nil,
	)

	if err != nil {
		log.Printf("Could not get Categories from the api")
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("User-agent", "Joke CLI (https://github.com/Raindevops/joke-cli)")

	rsp, err := http.DefaultClient.Do(r)

	if err != nil {
		log.Printf("Could not request the api")
	}

	responseBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("Could not read response Body. %v", err)
	}

	var cat []string

	if err := json.Unmarshal(responseBytes, &cat); err != nil {
		log.Printf("Could not unmarshall responsebyte. %v", err)
	}

	fmt.Printf("%v", strings.Join(cat, ", "))
}

func GetPhraseByCategory() {

}

// list categories
// joke by category
