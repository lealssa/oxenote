package main

import (
	"fmt"
	"log"
	"net/http"
	"oxenote/src/handlers"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlers.RootHandler)

	println("Start listening on 8080 port")
	err := http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
	if err != nil {
		fmt.Printf("Error starting application: %s", err.Error())
	}
}
