package main

import "math/rand"

var games []Game

func parse_into_db() {
	body := queryGame()
	games := readDataFromResponseBody(body)

	createDb()

	for _, game := range games {
		insertIntoDb(game.Name, game.Cover.ImageID)
	}
}

func pickRandomGame(games []Game) Game {
	return games[rand.Intn(len(games))]
}
