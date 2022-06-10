package controller

import (
	"go-web-practice/internal/config"
	"net/http"
)

func RegisterRoutes() {
	registerIndexRoutes()
	registerWelcomeRoutes()
	registerLookRoutes()
	http.NotFoundHandler()
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir(config.Config.Static))))
}
