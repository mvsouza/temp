package main

import (
	"fmt"
	"net/http"
)

// handler function for easier testing
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3001", nil)
}
