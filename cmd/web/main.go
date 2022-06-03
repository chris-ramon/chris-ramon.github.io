package main

import (
	"log"
	"net/http"
	"os"

	"embed"
)

//go:embed public/index.html public/style.css public/bg.jpg public/graphql.svg
var f embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := f.ReadFile("public/index.html")
		if err != nil {
			log.Printf("failed to find file: %v", err)
			w.Write([]byte("We'll back soon ..."))
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "https://use.fontawesome.com")
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write(b)
	})

	mux.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		b, err := f.ReadFile("public/style.css")
		if err != nil {
			log.Printf("failed to find file: %v", err)
			w.Write([]byte("We'll back soon ..."))
			return
		}

		w.Header().Add("Content-Type", "text/css")
		w.Write(b)
	})

	mux.HandleFunc("/bg.jpg", func(w http.ResponseWriter, r *http.Request) {
		b, err := f.ReadFile("public/bg.jpg")
		if err != nil {
			log.Printf("failed to find file: %v", err)
			w.Write([]byte("We'll back soon ..."))
			return
		}

		w.Header().Add("Content-Type", "image/png")
		w.Write(b)
	})

	mux.HandleFunc("/graphql.svg", func(w http.ResponseWriter, r *http.Request) {
		b, err := f.ReadFile("public/graphql.svg")
		if err != nil {
			log.Printf("failed to find file: %v", err)
			w.Write([]byte("We'll back soon ..."))
			return
		}

		w.Header().Add("Content-Type", "image/svg+xml")
		w.Write(b)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("server running on port :" + port)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
