package orgpa

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Query all the notes to the databaseAPI
// URL: /api/notes
// Write the databaseAPI's answer on the ResponseWriter
func (sh *ServerHandler) apiGetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	resp, err := http.Get(sh.Config.URLDatabaseAPI + "/list")
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{error: %s}", err.Error())
		return
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{error: %s}", err.Error())
		return
	}
	fmt.Fprint(w, string(responseBody))
}

func (sh *ServerHandler) apiNewNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		// To handle
		log.Println("error")
		return
	}
	// Do post request
}
