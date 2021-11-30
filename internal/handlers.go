package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)


func (s *server) getLanguages() http.HandlerFunc{
	return func(w http.ResponseWriter , r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.languages); err != nil{
			http.Error(w,err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) getLanguage() http.HandlerFunc{
	return func(w http.ResponseWriter, r*http.Request){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, item := range s.languages{
			if item.Code == params["code"]{
				if err := json.NewEncoder(w).Encode(item) ; err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}
}

func (s *server) translate() http.HandlerFunc{
	return func(w http.ResponseWriter, r*http.Request){
		response := WordResponse{TranslatedWord: "translatedWord" }
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}