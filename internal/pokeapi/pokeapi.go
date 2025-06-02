package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Fetches a LocationAreaResponse struct
func FetchLocRes(url string) (any, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get response body: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}

	var decodedRes LocationAreaResponse
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %v", err)
	}

	return &decodedRes, nil
}

// Fetches pokemon encountered on this location
func FetchEncounters(url string) (any, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get response body: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}

	var decodedRes EncounterResponse
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %v", err)
	}

	return &decodedRes, nil
}
