package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/onioncall/slugly"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    config, _ := slugly.NewHTTPSlugConfig(slugly.MediumPrintSpeed, w)
    config.SlugPrint("Hello! This is a test string... for testing.")
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server starting on port 8080...")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
