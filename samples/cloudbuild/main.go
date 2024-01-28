package main

import (
	"fmt"
	"log"
	"net/http"
)

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func handler(w http.ResponseWriter, r *http.Request) {
	greeting := getenv("GREETING", "Hello")
	fmt.Fprintln(w, "%s, world", greeting)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
