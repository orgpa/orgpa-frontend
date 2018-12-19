package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"orgpa-frontend/database"
	"strconv"

	"github.com/gorilla/mux"
)

func (sh ServerHandler) homePage(w http.ResponseWriter, r *http.Request) {
	err := sh.TmplEngine.GenerateAndExecuteTemplate(w, "HomePage", "orgpa - home", "")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (sh ServerHandler) notePage(w http.ResponseWriter, r *http.Request) {
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
		ID:       strconv.Itoa(note.ID),
		Content:  note.Content,
		Title:    note.Title,
		LastEdit: note.LastEdit,
	}

	sh.TmplEngine.GenerateAndExecuteTemplate(w, "NotePage", "orgpa - note", noteString)
}
