package main

import (
	"flag"
	"net/http"

	"github.com/catti_brie/examples/myserver/myhandler"
)

func main() {
	numPtr := flag.Int64("i", 0, "an int64 value v")
	portPtr := flag.String("p", ":8080", "port number, string value")
	flag.Parse()
	myhandler.V = *numPtr
	s := S{V: *numPtr}
	http.HandleFunc("/add", s.AddHandler)
	http.HandleFunc("/dec", myhandler.DecHandler)
	http.HandleFunc("/result", myhandler.ResHandler)
	http.HandleFunc("/set", myhandler.SetHandler)
	http.ListenAndServe(*portPtr, nil)
}
