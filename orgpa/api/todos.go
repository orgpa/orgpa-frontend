package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
