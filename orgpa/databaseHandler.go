package orgpa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"orgpa-frontend/database"

	"github.com/gorilla/mux"
)

// Query all the notes to the databaseAPI
// URL: /api/notes
// Write the databaseAPI's answer on the ResponseWriter
func (sh *ServerHandler) apiGetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	resp, err := http.Get(sh.Config.URLDatabaseAPI + "/list")
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": %s}", err.Error())
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": %s}", err.Error())
		return
	}
	fmt.Fprint(w, string(responseBody))
}

func (sh *ServerHandler) apiNewNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"missing information\"}")
		return
	}

	note := database.Notes{Title: title, Content: content}
	jsonData, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"'%s\"}", err.Error())
		return
	}

	resp, err := http.Post(sh.Config.URLDatabaseAPI+"/list", "applicaition/json", bytes.NewBuffer(jsonData))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
	}
	w.Write(body)
}

func (sh *ServerHandler) apiDeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"missing information\"}")
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", sh.Config.URLDatabaseAPI+"/"+id, nil)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}

	_, err = client.Do(req)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}
}
