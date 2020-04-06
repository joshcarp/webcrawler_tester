package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var testCases = map[string]string{
	"/1": "<a href=\"http://localhost/2\">",
	"/2": "<a href=\"localhost/3\">",
	"/3": "<a href=\"/4\">",
	"/4": "<a class=\"\" href=\"/1\">",
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	byts, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(byts))
	fmt.Println(r.URL.Path)

	fmt.Fprintf(w, "%s", testCases[r.URL.Path])
}
