package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("[Main] Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)

	err = godotenv.Load(".env", ".env.token")
	if err != nil {
		log.Fatalf("[Main] Error loading .env files: %v", err)
	}

	log.Println("[Main] Starting the application...")
}

func main() {
	/*	parse_into_db()*/
	/*server()*/
	server2()
}
