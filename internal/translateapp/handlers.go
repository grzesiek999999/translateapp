package translateapp

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
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

func (a *App) BatchTranslate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list []BatchTranslateResponse
		var words []WordToTranslate
		ctx := context.Background()
		err := r.ParseMultipartForm(1024)
		if err != nil {
			log.Fatal(err)
		}
		source := r.PostFormValue("source")
		target := r.PostFormValue("target")
		file, _, err := r.FormFile("text")
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					a.logger.Info("read file line error")
					break

				}
			}
			line = line[:len(line)-1]
			newWord := WordToTranslate{
				Source: source,
				Target: target,
				Word:   line,
			}
			words = append(words, newWord)
		}
		defer file.Close()

		for _, v := range words {
			response, _ := a.Service.BatchTranslate(ctx, v)
			list = append(list, *response)
		}
		json.NewEncoder(w).Encode(list)
		if err != nil {
			log.Fatal(err)
		}

	}

}
