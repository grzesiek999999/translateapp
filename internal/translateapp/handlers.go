package translateapp

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// maximum memory size of upload file in kb
const (
	maxMemory int64 = 1024
)

func (a *App) GetLanguages() http.HandlerFunc {
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

func (a *App) Translate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		ctx := context.Background()
		word := make([]WordToTranslate, 1, 1)
		err := json.NewDecoder(r.Body).Decode(&word[0])
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
		var wordsFile []string
		var wordsTranslate []WordToTranslate
		ctx := context.Background()
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			a.logger.Error("File is too big, maximum file size is 1024kb")
		}

		source := r.PostFormValue("source")
		target := r.PostFormValue("target")
		file, _, err := r.FormFile("text")

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			wordsFile = append(wordsFile, scanner.Text())
		}

		for _, v := range wordsFile {
			newWord := WordToTranslate{
				Source: source,
				Target: target,
				Word:   v,
			}
			wordsTranslate = append(wordsTranslate, newWord)
		}

		defer file.Close()

		response, _ := a.Service.Translate(ctx, wordsTranslate)
		json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Fatal(err)
		}

	}

}
