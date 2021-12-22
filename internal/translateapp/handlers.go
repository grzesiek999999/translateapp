package translateapp

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (a *App) getLanguages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		response, err := a.Service.GetLanguages(ctx)
		if err != nil {
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (a *App) translate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		var word WordToTranslate
		err := json.NewDecoder(r.Body).Decode(&word)
		if err != nil {
			log.Fatal(err)
		}
		response, err := a.Service.Translate(ctx, word)
		if err != nil {
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
