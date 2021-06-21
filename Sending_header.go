package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A GO web Server")
	w.WriteHeader(200)
	fmt.Println("Anil")

}
