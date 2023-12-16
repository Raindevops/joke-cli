package api

import (
	"io"
	"log"
	"net/http"
)

func GetDataFromApi(uri string) []byte {
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
