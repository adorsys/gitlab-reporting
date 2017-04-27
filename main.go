package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// port := flag.String("port", "9090", "set the port where the tool shall run")
	// ip := flag.String("ip", "127.0.0.1", "set the ip adress where the tool shall run")
	// flag.Parse()

	// Supplied by CloudFoundry
	port := os.Getenv("PORT")

	fmt.Printf("Start reporting tool")

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/createReport", createReport)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server at: :%v", port)
	server.ListenAndServe()
}
