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
	//myhandler.V = *numPtr
	var i myhandler.Operat = &myhandler.Value{V: *numPtr}
	var h = myhandler.Handlers{H: i}
	http.HandleFunc("/add", h.AddHandler)
	http.HandleFunc("/dec", h.DecHandler)
	http.HandleFunc("/result", h.ResHandler)
	http.HandleFunc("/set", h.SetHandler)
	http.ListenAndServe(*portPtr, nil)
}
