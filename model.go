package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Cover struct {
	ID      int    `json:"id"`
	ImageID string `json:"image_id"`
}

type Platform struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Game struct {
	ID        int        `json:"id"`
	Cover     Cover      `json:"cover"`
	Name      string     `json:"name"`
	Platforms []Platform `json:"platforms"`
	Rating    float64    `json:"rating"`
	Summary   string     `json:"summary"`
}

func readData(filename string) []Game {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var games []Game
	if err := json.Unmarshal(data, &games); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return games
}

func displayGameDetails(games []Game) {
	for _, game := range games {
		fmt.Printf("Game Name: %s\n", game.Name)
		fmt.Printf("Rating: %.2f\n", game.Rating)
		fmt.Printf("Summary: %s\n", game.Summary)
		fmt.Println("Platforms:")
		for _, platform := range game.Platforms {
			fmt.Printf("- %s\n", platform.Name)
		}
		fmt.Println()
	}
}
