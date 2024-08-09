package handlers

import (
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/spa.html", "templates/_note.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sua lógica específica para a rota raiz aqui...

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
