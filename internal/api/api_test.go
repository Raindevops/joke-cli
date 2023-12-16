package api

import "testing"

func TestGetDataFromAPIWithGoodURL(t *testing.T) {
	const uri = "https://api.chucknorris.io/jokes/random"

	bytes := GetDataFromApi(uri)

	if len(bytes) <= 0 {
		t.Errorf("The value returned by the api is not an array of bytes")
	}
}
