package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type pageData struct {
	APIKey string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env error")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			http.Error(w, "API_KEY not set", http.StatusInternalServerError)
			return
		}

		data := pageData{APIKey: apiKey}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
