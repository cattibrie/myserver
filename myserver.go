package main

import (
	"fmt"
	"net/http"
)

var v int

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		v++
		fmt.Fprintf(w, "Value v = %d", v)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func decHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		v--
		fmt.Fprintf(w, "Value v = %d", v)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func resHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Value v = %d", v)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/dec", decHandler)
	http.HandleFunc("/result", resHandler)
	http.ListenAndServe(":8080", nil)
}
