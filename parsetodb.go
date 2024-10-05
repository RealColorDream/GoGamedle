package main

func parse_into_db() {
	body := queryGame()
	games := readDataFromResponseBody(body)

	createDb()

	for _, game := range games {
		insertIntoDb(game.Name, game.Cover.ImageID)
	}
}
