package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"orgpa-frontend/database"

	"github.com/gorilla/mux"
)

func (apiH *Handler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	resp, err := http.Get(apiH.URLDatabaseAPI + "/notes")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}
	fmt.Fprint(w, string(responseBody))
}

func (apiH *Handler) newNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	// Get title and content value
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "missing information"}`)
		return
	}

	// Create the new note and tramsform in JSON
	note := database.Notes{Title: title, Content: content}
	jsonData, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	// Send the new note to the database API
	resp, err := http.Post(apiH.URLDatabaseAPI+"/notes", "applicaition/json", bytes.NewBuffer(jsonData))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
	}

	// Return the new note in JSON
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
	}
	w.Write(body)
}

func (apiH *Handler) deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf8")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "missing information"}`)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", apiH.URLDatabaseAPI+"/notes/"+id, nil)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
	}
	w.Write(body)
}

func (apiH *Handler) patchNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Get the note's ID and content
	id := r.FormValue("id")
	content := r.FormValue("content")
	if id == "" || content == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "missing information"}`)
		return
	}

	// Create a JSON of the note
	note := database.Notes{Content: content}
	jsonData, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	// Create the request
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", apiH.URLDatabaseAPI+"/notes/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	// Return the patched note in JSON
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
	}
	w.Write(body)
}
