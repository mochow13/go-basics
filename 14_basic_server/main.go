package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Angular Reaction</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>All these frameworks</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server launced...")
	http.HandleFunc("/about", about)
	http.ListenAndServe(":3000", nil)
}
