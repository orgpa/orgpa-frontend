package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (apiH *Handler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	resp, err := http.Get(apiH.URLDatabaseAPI + "/todos")
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

func (apiH *Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
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
	req, err := http.NewRequest("DELETE", apiH.URLDatabaseAPI+"/todos/"+id, nil)
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
