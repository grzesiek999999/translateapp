package main

import (
	"fmt"
	"log"
	"os"
	app "translateapp/internal/app/translateapp"
	"translateapp/internal/logging"
)

func main() {
	logger := logging.NewLogger("debug", true).Desugar()
	
	if err := app.Run(logger); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		log.Fatal(err)
	}
}
