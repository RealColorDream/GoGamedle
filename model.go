package main

import (
	"encoding/json"
	"log"
)

type Cover struct {
	ID      int    `json:"id"`
	ImageID string `json:"image_id"`
}

type Game struct {
	Name  string `json:"name"`
	Cover Cover  `json:"cover"`
}

func readDataFromResponseBody(responseBody []byte) []Game {
	var games []Game
	if err := json.Unmarshal(responseBody, &games); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return games
}
