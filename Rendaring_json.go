package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	FirstName string
	Lastname  string
	Hobbies   []string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	Profile := Profile{"Anil", "kumar", []string{
		"music", "dance", "singing"}}

	js, err := json.Marshal(Profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
