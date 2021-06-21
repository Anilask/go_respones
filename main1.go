package main

import (
	"net/http"
	"path"
	"text/template"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"music", "dancing"}}

	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index1.html")

	// note that the layout file be the first parmeters in the parserfile
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
}
