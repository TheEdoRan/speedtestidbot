package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/theedoran/speedtestidbot/bot"
	"github.com/theedoran/speedtestidbot/fly"
)

func main() {
	// Use godotenv only if .env exists in current directory.
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("error: could not load environmental vars")
		}
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go fly.HealthCheck()
	fmt.Println("healthcheck server started")

	bot.Start(done)
	fmt.Println("bot gracefully stopped")
}
