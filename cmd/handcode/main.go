package main

import (
	"log"
	"os"

	"github.com/slackerkids/handcode/internal/app"
)

func main() {
	application := app.New()

	if err := application.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
