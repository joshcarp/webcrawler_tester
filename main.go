package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type Server struct {
	http.Handler
	verbose bool
}


func main() {
	s := Server{Handler: http.FileServer(http.Dir("./web1.comp30023"))}

	// Parse verbose mode
	flag.BoolVar(&s.verbose, "v", false, "Verbose mode: Prints out request")
	flag.Parse()

	// Start up a server
	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", s)
	if err != nil {
		log.Fatal(err)
	}
}

//ServeHTTP writes out debug info to stdout and redirects the request to the filesystem handler
func (s Server)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Visited URL: %s\n",r.URL)
	if s.verbose{
		byts, _ := httputil.DumpRequest(r, true)
		fmt.Printf("Request sent: %s\n\n",string(byts))
	}
	// Forward the response onto the filesystem handler
	s.Handler.ServeHTTP(w,r)
}
