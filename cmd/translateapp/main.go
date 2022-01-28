package main

import (
	"log"
	app "translateapp/internal/app/translateapp"
	"translateapp/internal/logging"
)

func main() {
	logger := logging.NewLogger("debug", true).Desugar()
	if err := app.Run(logger); err != nil {
		log.Fatal(err)
	}
}
