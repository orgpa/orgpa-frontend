package orgpa

import (
	"html/template"
	"orgpa-frontend/configuration"

	"github.com/gorilla/mux"
)

// ServerHandler contains the template pointer and the
// service configuration and all the routes.
type ServerHandler struct {
	Template *template.Template
	Config   configuration.ServiceConfig
}

// Return a new ServerHandler with the given configuration
// and Template pointer.
func newServerHandler(config configuration.ServiceConfig, template *template.Template) ServerHandler {
	return ServerHandler{
		Template: template,
		Config:   config,
	}
}

// Define all the ServerHandler route to the given Router.
func (sh *ServerHandler) defineRoutes(router *mux.Router) {
	// Main route
	router.Methods("GET").Path("/").HandlerFunc(sh.homePage)

	// Notes route
	notesRouter := router.PathPrefix("/note").Subrouter()
	notesRouter.Methods("GET").Path("/{id}").HandlerFunc(sh.notePage)

	// Static file route
	router.Methods("GET").PathPrefix("/static/").HandlerFunc(sh.serveStatic)

	// API route
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Methods("GET").Path("/notes").HandlerFunc(sh.apiGetAllNotes)
	apiRouter.Methods("POST").Path("/notes").HandlerFunc(sh.apiNewNote)
	apiRouter.Methods("PATCH").Path("/notes").HandlerFunc(sh.apiPatchNote)
	apiRouter.Methods("DELETE").Path("/notes/{id}").HandlerFunc(sh.apiDeleteNote)
}
