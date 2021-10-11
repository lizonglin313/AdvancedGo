package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	h := handleRequest("456")
	http.HandleFunc("/1", h)
	server.ListenAndServe()
}

func handleRequest(t string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		get(w, r, t)
		fmt.Println(t)
	}
}

func get(w http.ResponseWriter, r *http.Request, t string)  {
	fmt.Fprintf(w, t)
}
