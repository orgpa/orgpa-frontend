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
	router.Methods("GET").Path("/").HandlerFunc(sh.homePage)

	router.Methods("GET").PathPrefix("/static/").HandlerFunc(sh.serveStatic)
}
