package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	next     *string
	previous *string
}

func MapNextLocationAreas(config *Config) error {
	if config.next == nil {
		rootUrl := "https://pokeapi.co/api/v2/location-area"
		config.next = &rootUrl
	}
	err, la := mapLocationAreas(*config.next)
	if err != nil {
		return err
	}
	for _, v := range la.Results {
		fmt.Println(v.Name)
	}

	config.previous = la.Previous
	config.next = la.Next
	return nil
}

func MapPreviousLocationAreas(config *Config) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	err, la := mapLocationAreas(*config.previous)
	if err != nil {
		return err
	}
	for _, v := range la.Results {
		fmt.Println(v.Name)
	}
	
	config.previous = la.Previous
	config.next = la.Next
	return nil
}

func mapLocationAreas(url string) (error, LocationAreas) {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error when fetching location-area %v", err), LocationAreas{}
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body), LocationAreas{}
	}
	if err != nil {
		return fmt.Errorf("error when reading body %v", err), LocationAreas{}
	}
	la := LocationAreas{}
	err = json.Unmarshal(body, &la)
	if err != nil {
		return fmt.Errorf("error when unmarshaling %v", err), LocationAreas{}
	}

	return nil, la
}

type LocationAreas struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
