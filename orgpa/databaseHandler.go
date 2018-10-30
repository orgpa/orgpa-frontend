package orgpa

import (
	"fmt"
	"net/http"
)

// Query all the notes to the databaseAPI
// URL: /api/notes
// Write the databaseAPI's answer on the ResponseWriter
func (sh *ServerHandler) apiGetAllNotes(w http.ResponseWriter, r *http.Request) {

	// token, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "{error: %s}", err.Error())
	// 	return
	// }
	fmt.Fprint(w, r)
	fmt.Println(r)
}
