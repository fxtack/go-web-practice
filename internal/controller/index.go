package controller

import (
	"go-web-practic/internal/config"
	"go-web-practic/internal/templates"
	"net/http"
	"strconv"
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
			"pprof":   strconv.Itoa(config.Config.PprofPort),
		})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
