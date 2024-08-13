package main

import (
	"fmt"
	"log"
	"net/http"
	"oxenote/src/auth"
	"oxenote/src/handlers"

	"github.com/gorilla/sessions"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

var store = sessions.NewCookieStore([]byte("sua-chave-secreta"))

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica se o usuário está autenticado
		if !auth.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// Se o usuário estiver autenticado, permite o acesso à próxima função
		next.ServeHTTP(w, r)
	})
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		AuthRequired(http.HandlerFunc(handlers.RootHandler)).ServeHTTP(w, r)
	})

	http.HandleFunc("/login", handlers.LoginHandler)

	println("Start listening on 8080 port")
	err := http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
	if err != nil {
		fmt.Printf("Error starting application: %s", err.Error())
	}
}
