package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux" // Trying Gorilla Mux
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)

	log.Printf("Serving on port 8080")

	err := http.ListenAndServe(":8080", r)

	log.Fatal(err)

}
