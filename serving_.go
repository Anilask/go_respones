package main

import (
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	// Assuming tou want to server a photo at 'images/foo.png'
	fp := path.Join("imagees", "foo.png")
	http.ServeFile(w, r, fp)
}
