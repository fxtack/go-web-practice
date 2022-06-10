package controller

import (
	"go-web-practic/internal/templates"
	"net/http"
	"strings"
)

func registerLookRoutes() {
	http.HandleFunc("/look/", handleLook)
}

func handleLook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 切割 URL 来获取 URL 中的参数
		p := strings.Split(r.URL.Path, "/")
		switch p[2] {
		case "go":
			w.Header().Set("Location", "https://studygolang.com/pkgdoc")
			w.WriteHeader(302)
		case "down":
			templates.Template.Lookup("look.tmpl").Execute(w, map[string]string{
				"title": "Look down",
				"image": "golang-down.png",
			})
		case "":
			templates.Template.Lookup("look.tmpl").Execute(w, map[string]string{
				"title": "Look right",
				"image": "golang-right.png",
			})
		default:
			w.WriteHeader(http.StatusNotFound)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
