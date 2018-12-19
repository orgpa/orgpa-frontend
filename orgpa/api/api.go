package api

import (
	"orgpa-frontend/configuration"
)

// Handler is the main struct that allows to handle all the
// request on the API.
type Handler struct {
	URLDatabaseAPI string
}

// NewHandler returns a new handler with the given configuration.
func NewHandler(config configuration.ServiceConfig) Handler {
	return Handler{
		URLDatabaseAPI: config.URLDatabaseAPI,
	}
}
