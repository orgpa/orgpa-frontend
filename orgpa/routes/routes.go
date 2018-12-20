package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (sh ServerHandler) homePage(w http.ResponseWriter, r *http.Request) {
	err := sh.TmplEngine.GenerateAndExecuteTemplate(w, "HomePage", "orgpa - home", "")
	// debug
	if err != nil {
		log.Println(err.Error())
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
	resp, err := http.Get(sh.Config.URLDatabaseAPI + "/notes/" + varID)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", 400)
		return
	}

	// Get JSON from response's body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", 400)
		return
	}

	// Decode JSON
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", 400)
		return
	}
	itemsMap := f.(map[string]interface{})

	err = sh.TmplEngine.GenerateAndExecuteTemplate(w, "NotePage", "orgpa - note", itemsMap)
	// debug
	if err != nil {
		log.Println(err.Error())
	}
}
