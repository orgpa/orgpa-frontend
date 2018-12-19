package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"orgpa-frontend/database"
	"time"

	"github.com/gorilla/mux"
)

func (apiH *Handler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	resp, err := http.Get(apiH.URLDatabaseAPI + "/list")
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

func (apiH *Handler) newNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"missing information\"}")
		return
	}

	note := database.Notes{ID: 0, Title: title, Content: content, LastEdit: time.Now().UTC()}
	jsonData, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"'%s\"}", err.Error())
		return
	}

	resp, err := http.Post(apiH.URLDatabaseAPI+"/list", "applicaition/json", bytes.NewBuffer(jsonData))
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

func (apiH *Handler) deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf8")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"missing information\"}")
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", apiH.URLDatabaseAPI+"/list/"+id, nil)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}

	_, err = client.Do(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}
}

func (apiH *Handler) patchNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf8")
	id := r.FormValue("id")
	content := r.FormValue("content")
	if id == "" || content == "" {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"missing information\"}")
		return
	}

	note := database.Notes{Content: content}
	jsonData, err := json.Marshal(note)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"'%s\"}", err.Error())
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", apiH.URLDatabaseAPI+"/list/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}

	_, err = client.Do(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		return
	}
}
