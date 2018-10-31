package orgpa

import (
	"fmt"
	"io/ioutil"
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
