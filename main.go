package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type Server struct {
	Hosts map[string]http.Handler
	Verbose bool
}

func main() {

	// Parse Verbose mode
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Verbose mode: Prints out request")
	flag.Parse()

	serveHost(verbose, "web1.comp30023", "web2.comp30023", "web3.comp30023")
}

func serveHost(v bool, hosts ...string ){
	// Start up a server
	s := Server{Verbose: v}
	s.Hosts = make(map[string]http.Handler)
	for _, host := range hosts{
		s.Hosts[host] = http.FileServer(http.Dir(host))
	}
	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", s)
	if err != nil {
		log.Fatal(err)
	}
}

//ServeHTTP writes out debug info to stdout and redirects the request to the filesystem handler
func (s Server)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Visited URL: %s\n",r.URL)
	if s.Verbose {
		byts, _ := httputil.DumpRequest(r, true)
		fmt.Printf("Request sent: %s\n\n",string(byts))
	}
	// Forward the response onto the filesystem handler
	println(r.Host)
	s.Hosts[r.Host].ServeHTTP(w,r)
}
