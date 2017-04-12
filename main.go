package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/createReport", createReport)
	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: mux,
	}
	server.ListenAndServe()
}
