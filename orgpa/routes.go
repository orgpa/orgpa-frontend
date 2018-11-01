package orgpa

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func (sh *ServerHandler) homePage(w http.ResponseWriter, r *http.Request) {

	// Usage of template for the tests.
	t, _ := template.ParseFiles("./frontend/views/HomePage.html")
	t.Execute(w, sh)

	// sh.Template.Execute(w, "HomePage.html", sh)
}

// Serve static filese
func (sh *ServerHandler) serveStatic(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[len("/static/"):]

	if strings.HasSuffix(sh.Config.StaticFilePath, "/") == false {
		sh.Config.StaticFilePath += "/"
	}
	path = sh.Config.StaticFilePath + path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".jpg") {
			contentType = "image/jpg"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		w.Write([]byte("{error:" + err.Error() + "}"))
	}
}
