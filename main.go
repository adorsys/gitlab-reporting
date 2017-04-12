package main

import (
	"flag"
	"net/http"
)

func main() {
	port := flag.String("port", "9090", "set the port where the tool shall run")
	flag.Parse()

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/createReport", createReport)
	server := &http.Server{
		Addr:    "127.0.0.1:" + *port,
		Handler: mux,
	}
	server.ListenAndServe()
}
