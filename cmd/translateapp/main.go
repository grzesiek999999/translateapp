package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"translateapp/internal"
)

func main() {
	log.Printf("starting...")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	server := internal.NewServer()
	http.ListenAndServe(":8080", server)
	defer done()
	<-ctx.Done()
	log.Printf("successful shutdown")
}