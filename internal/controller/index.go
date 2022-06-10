package controller

import (
	"go-web-practice/internal/templates"
	"net/http"
)

func registerIndexRoutes() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates.Template.Lookup("index.tmpl").Execute(w, map[string]string{
			"welcome": "/welcome",
			"look":    "/look",
			"statics": "/web",
		})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
