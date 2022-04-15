package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/sayHello", sayHello).Methods("GET", "POST")
	http.ListenAndServe(":3009", r)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello!!"))
	fmt.Fprint(w, "Hello!!")
}
