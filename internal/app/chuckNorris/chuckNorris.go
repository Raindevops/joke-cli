package chucknorris

import (
	"encoding/json"
	"fmt"
	"joke-cli/internal/api"
	"log"
	"strings"
)

const api_url = "https://api.chucknorris.io/jokes/"

type Phrase struct {
	ID     string `json:"id"`
	Phrase string `json:"value"`
}

func RandomPhrase() {
	responseBytes := api.GetDataFromApi(api_url + "random")
	phrase := Phrase{}
	if err := json.Unmarshal(responseBytes, &phrase); err != nil {
		fmt.Printf("Could not unmarshal responseByte. %v", err)
	}

	fmt.Printf(phrase.Phrase)
}

func ListAllCategories() {
	rspBytes := api.GetDataFromApi(api_url + "categories")
	var cat []string

	if err := json.Unmarshal(rspBytes, &cat); err != nil {
		log.Printf("Could not unmarshall responseByte. %v", err)
	}

	fmt.Printf("%v", strings.Join(cat, ", "))
}

func GetPhraseByCategory(category string) {
	rspBytes := api.GetDataFromApi(api_url + "random?category=" + category)
	phrase := Phrase{}

	if err := json.Unmarshal(rspBytes, &phrase); err != nil {
		log.Printf("Could not unmarshall responseBytes. %v", err)
	}
	fmt.Printf(phrase.Phrase)
}
