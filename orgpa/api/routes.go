package api

import "github.com/gorilla/mux"

func (apiH *Handler) DefineRoute(r *mux.Router) {
	// base url for API
	apiSubrouter := r.PathPrefix("/api").Subrouter()

	// API notes URL
	apiSubrouter.Methods("GET").Path("/notes").HandlerFunc(apiH.getAllNotes)
	apiSubrouter.Methods("POST").Path("/notes").HandlerFunc(apiH.newNote)
	apiSubrouter.Methods("PATCH").Path("/notes").HandlerFunc(apiH.patchNote)
	apiSubrouter.Methods("DELETE").Path("/notes/{id}").HandlerFunc(apiH.deleteNote)

	// API todos URL
	apiSubrouter.Methods("GET").Path("/todos").HandlerFunc(apiH.getAllTodos)
}
