package main

import (
	"encoding/xml"
	"net/http"
)

type Profile struct {
	FirstName string
	LastName  string
	Hobbies   []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Anil", "kumar", []string{"music", "dance", "cricket"}}

	x, err := xml.MarshalIndent(profile, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
