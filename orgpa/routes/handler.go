package routes

import (
	"orgpa-frontend/configuration"
	"orgpa-frontend/template"

	"github.com/gorilla/mux"
)

type ServerHandler struct {
	TmplEngine template.TemplateEngine
	Config     configuration.ServiceConfig
}

// NewHandler returns a new handler with the given config and template engine
func NewHandler(config configuration.ServiceConfig, templateEngine template.TemplateEngine) ServerHandler {
	return ServerHandler{
		TmplEngine: templateEngine,
		Config:     config,
	}
}

// DefineMainRoute defines the main routes of the server
func (sh ServerHandler) DefineMainRoute(r *mux.Router) {
	// Main route
	r.Methods("GET").Path("/").HandlerFunc(sh.homePage)

	// Notes route
	notesSubrouter := r.PathPrefix("/note").Subrouter()
	notesSubrouter.Methods("GET").Path("/{id}").HandlerFunc(sh.notePage)
}

// DefineStaticRoute defines the route to server static files
func (sh ServerHandler) DefineStaticRoute(r *mux.Router) {
	r.Methods("GET").PathPrefix("/static/").HandlerFunc(sh.serveStatic)
}
