package orgpa

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"orgpa-frontend/database"

	"github.com/gorilla/mux"
)

func (sh *ServerHandler) homePage(w http.ResponseWriter, r *http.Request) {
	pi := newPageInfo("orgpa - home", "", sh.Config)
	t, _ := template.ParseFiles("./frontend/views/HomePage.html")
	t.Execute(w, pi)
}

func (sh *ServerHandler) notePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Get ID from arguments
	varID, ok := vars["id"]
	if !ok {
		http.Redirect(w, r, "/", 400)
		return
	}

	// Request database API
	resp, err := http.Get(sh.Config.URLDatabaseAPI + "/list/" + varID)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", 400)
		return
	}

	// Decode JSON from response's body
	var note database.Notes
	err = json.NewDecoder(resp.Body).Decode(&note)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", 400)
		return
	}

	// Create a renderable Note structure from the previous
	// requested note struct.
	var noteString = database.NotesString{
		ID:       note.ID.Hex(),
		Content:  note.Content,
		Title:    note.Title,
		LastEdit: note.LastEdit,
	}

	pi := newPageInfo("orgpa - note", noteString, sh.Config)
	t, _ := template.ParseFiles("./frontend/views/NotePage.html")
	t.Execute(w, pi)
}
