package main

import (
	"flag"
	"log"
	"net/http"
)

var urls = make(map[string]string)

func home(w http.ResponseWriter, r *http.Request) {
	urls["//test.co"] = "//google.com"
	urls["//test2.co"] = "//bbc.co.uk"
}

func main() {
	addr := flag.String("addr", ":9990", "HTTP network address")
	flag.Parse()

	http.HandleFunc("/", home)
	log.Printf("Starting server on http://localhost%s:\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
