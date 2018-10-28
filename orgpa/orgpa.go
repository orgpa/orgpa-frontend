package orgpa

import (
	"html/template"
	"net/http"
	"orgpa-frontend/configuration"
	"time"

	"github.com/gorilla/mux"
)

// Run the frontend Orgpa server
func Run(config configuration.ServiceConfig, template *template.Template) error {
	handler := newServerHandler(config, template)
	router := mux.NewRouter()

	handler.defineRoutes(router)

	srv := http.Server{
		Addr:           config.Endpoint,
		Handler:        router,
		IdleTimeout:    5 * time.Second,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv.ListenAndServe()
}
