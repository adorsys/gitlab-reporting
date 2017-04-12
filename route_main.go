package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "sidebar", "report")
}

func createReport(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	url := r.Form.Get("url")
	token := r.Form.Get("token")
	branch := r.Form.Get("branch")
	project := r.Form.Get("project")
	apiversion := r.Form.Get("apiversion")

	if apiversion == "" {
		apiversion = "3"
	}

	if branch == "" {
		branch = ""
	}

	report, err := createReportData(url, token, branch, project, apiversion)

	if err != nil {
		fmt.Printf("createReport(): Error: %v", err)
	} else {
		generateHTML(w, report, "layout", "sidebar", "report")
	}
}
