package orgpa

import (
	"log"
	"net/http"
	"orgpa-frontend/orgpa/api"
	"time"

	"orgpa-frontend/configuration"
	"orgpa-frontend/template"

	"github.com/gorilla/mux"
)

// Run the frontend Orgpa server
func Run(config configuration.ServiceConfig) error {
	tmplEngine := template.NewTemplateEngine(config)
	handler := newServerHandler(config, tmplEngine)
	apiHandler := api.NewHandler(config)
	r := mux.NewRouter()

	handler.defineMainRoute(r)
	handler.defineStaticRoute(r)
	apiHandler.DefineRoute(r)

	log.Println("run")

	srv := http.Server{
		Addr:           config.Endpoint,
		Handler:        r,
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv.ListenAndServe()
}

// Define the main routes of the server
func (sh *ServerHandler) defineMainRoute(r *mux.Router) {
	// Main route
	r.Methods("GET").Path("/").HandlerFunc(sh.homePage)

	// Notes route
	notesSubrouter := r.PathPrefix("/note").Subrouter()
	notesSubrouter.Methods("GET").Path("/{id}").HandlerFunc(sh.notePage)
}

// Define the route to server static files
func (sh *ServerHandler) defineStaticRoute(r *mux.Router) {
	r.Methods("GET").PathPrefix("/static/").HandlerFunc(sh.serveStatic)
}
