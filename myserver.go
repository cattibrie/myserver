package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
)

var v int64

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nv := atomic.AddInt64(&v, 1)
		fmt.Fprintf(w, "Value v = %d", nv)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func decHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nv := atomic.AddInt64(&v, -1)
		fmt.Fprintf(w, "Value v = %d", nv)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func resHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		nv := atomic.LoadInt64(&v)
		fmt.Fprintf(w, "Value v = %d", nv)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sval := r.URL.Query().Get("value")
		val, err := strconv.ParseInt(sval, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid value argument: %v", err)
			return
		}
		atomic.StoreInt64(&v, val)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/dec", decHandler)
	http.HandleFunc("/result", resHandler)
	http.HandleFunc("/set", setHandler)
	http.ListenAndServe(":8080", nil)
}
