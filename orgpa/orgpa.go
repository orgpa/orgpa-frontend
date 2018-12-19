package orgpa

import (
	"log"
	"net/http"
	"orgpa-frontend/orgpa/api"
	"orgpa-frontend/orgpa/routes"
	"time"

	"orgpa-frontend/configuration"
	"orgpa-frontend/template"

	"github.com/gorilla/mux"
)

// Run the frontend Orgpa server
func Run(config configuration.ServiceConfig) error {
	tmplEngine := template.NewTemplateEngine(config)
	handler := routes.NewHandler(config, tmplEngine)
	apiHandler := api.NewHandler(config)
	r := mux.NewRouter()

	handler.DefineMainRoute(r)
	handler.DefineStaticRoute(r)
	apiHandler.DefineRoute(r)

	srv := http.Server{
		Addr:           config.Endpoint,
		Handler:        r,
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server running on :", config.Endpoint)

	return srv.ListenAndServe()
}
