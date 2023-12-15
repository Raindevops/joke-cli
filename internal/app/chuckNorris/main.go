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
	responseBytes := getDataFromApi(api_url + "random")
	phrase := Phrase{}
	if err := json.Unmarshal(responseBytes, &phrase); err != nil {
		fmt.Printf("Could not unmarshal responseByte. %v", err)
	}

	fmt.Println(string(phrase.Phrase))
}

func ListAllCategories() {
	rspBytes := getDataFromApi(api_url + "categories")
	var cat []string

	if err := json.Unmarshal(rspBytes, &cat); err != nil {
		log.Printf("Could not unmarshall responseByte. %v", err)
	}

	fmt.Printf("%v", strings.Join(cat, ", "))
}

func GetPhraseByCategory(category string) {
	rspBytes := getDataFromApi(api_url + "random?category=" + category)
	phrase := Phrase{}

	if err := json.Unmarshal(rspBytes, &phrase); err != nil {
		log.Printf("Could not unmarshall responseBytes. %v", err)
	}

	fmt.Printf(phrase.Phrase)

}

func getDataFromApi(url string) []byte {
	r, err := http.NewRequest(
		http.MethodGet,
		url,
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
