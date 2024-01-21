package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
