package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("port", "9090", "set the port where the tool shall run")
	ip := flag.String("ip", "127.0.0.1", "set the ip adress where the tool shall run")
	flag.Parse()

	fmt.Printf("Start reporting tool")

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/createReport", createReport)
	server := &http.Server{
		Addr:    *ip + ":" + *port,
		Handler: mux,
	}

	fmt.Printf("Starting server at: %v:%v", *ip, *port)
	server.ListenAndServe()
}
