package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/stephdewit/wordle-solver/solver"
)

const length = 5

var cachedWords []solver.Word

func main() {
	enableCORS := flag.Bool("cors", false, "enable CORS middleware")
	flag.Parse()

	var err error
	cachedWords, err = solver.LoadWords(length, false)
	if err != nil {
		log.Fatalf("failed to load word list: %v", err)
	}
	log.Printf("loaded %d words", len(cachedWords))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := http.HandlerFunc(handleSuggest)
	if *enableCORS {
		handler = cors(handleSuggest)
		log.Println("CORS enabled")
	}
	http.Handle("/suggestions", handler)

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
