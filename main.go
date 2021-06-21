// html file 
package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path"
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

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)

	/*if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}*/
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, profile); err != nil {
		http.Error(w, err.Error().Http.StatusInternalServerError)
	}
	templateString := buf.String()
}
