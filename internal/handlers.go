package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func (s *Server) getLanguages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		sugar := logger.Sugar()
		sugar.Infof("GET request on localhost:8080/languages")
		if err := json.NewEncoder(w).Encode(s.languages); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getLanguage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, item := range s.languages {
			if item.Code == params["code"] {
				if err := json.NewEncoder(w).Encode(item); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}
}

func (s *Server) translate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		response := WordResponse{TranslatedWord: "translatedWord"}
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		sugar := logger.Sugar()
		sugar.Infof("GET request on localhost:8080/translate")
		if logger == nil {
			log.Println("request type: GET, endpoint: localhost:8080/languages")
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
