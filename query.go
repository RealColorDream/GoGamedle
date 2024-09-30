package main

import (
	"github.com/Henry-Sarabia/apicalypse"
	"io"
	"log"
	"net/http"
	"os"
)

func queryGame() {
	clientID := os.Getenv("CLIENT_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")

	opt := apicalypse.ComposeOptions(
		apicalypse.Fields("name, cover.image_id, parent_game, storyline, rating, summary, platforms.name"),
		apicalypse.Where("genres.name = \"Shooter\""),
		apicalypse.Sort("rating", "desc"),
		apicalypse.Limit(500),
	)

	req, err := apicalypse.NewRequest("POST",
		"https://api.igdb.com/v4/games",
		opt,
	)

	if err != nil {
		log.Fatalf("[QUERY] Failed to create request: %v", err)
	}

	req.Header.Set("Client-ID", clientID)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	var client = &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[QUERY] Failed to read response body: %v", err)
	}

	err = os.WriteFile("response_body.json", body, 0644)

	if err != nil {
		log.Fatalf("[QUERY] Failed to write response body to file: %v", err)
	}

	log.Printf("[QUERY] Response body saved to response_body.json")
}
