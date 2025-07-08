package pokeapi

import (
	"io"
	"log"
	"net/http"
)

func CallAPI() ([]byte, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // always close the body when done

	body, err := io.ReadAll(res.Body) // read the body
	if err != nil {
		return nil, err
	}
	return body, nil // return the raw data
}
