package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/woelkjulian/gitlab-reporting/gitlab"
	"github.com/woelkjulian/gitlab-reporting/model"
)

func generateHTML(w http.ResponseWriter, data interface{}, fileNames ...string) {
	var files []string
	for _, file := range fileNames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	funcMap := template.FuncMap{"fdate": formatDate}
	templates := template.New("templates/report.html").Funcs(funcMap)
	templates = template.Must(templates.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func formatDate(date string) (string, error) {
	time, err := time.Parse(time.RFC3339, date)

	if err != nil {
		return "", err
	}

	strTime := fmt.Sprintf("%v.%v.%v %v:%v Uhr", time.Day(), time.Month(), time.Year(), time.Hour(), time.Minute())
	return strTime, nil
}

func createReportData(url, token, branch, project, apiversion string) (model.Report, error) {
	var report model.Report
	var mergeRequestCommits []model.Commit
	var mergeRequestsWithChanges []model.MergeRequestWithChanges

	if apiversion == "" {
		report.Project.ApiVersion = "3"
	}

	commits, err := gitlab.GetAllCommits(&url, &token, &branch, &project, &apiversion, 0)
	if err != nil {
		return report, err
	}
	mergeRequests, err := gitlab.GetMergeRequests(&url, &token, &branch, &project, &apiversion, 0)
	if err != nil {
		return report, err
	}

	for _, req := range mergeRequests {
		newCommits, err := gitlab.GetMergeRequestCommits(&url, &token, &branch, &project, &apiversion, req.Id, 0)
		if err != nil {
			return report, err
		}
		newChanges, err := gitlab.GetMergeRequestChanges(&url, &token, &branch, &project, &apiversion, req.Id, 0)
		if err != nil {
			return report, err
		}
		for _, c := range newCommits {
			mergeRequestCommits = append(mergeRequestCommits, c)
		}
		mergeRequestsWithChanges = append(mergeRequestsWithChanges, newChanges)
	}

	report.Project.Url = url
	report.Project.Token = token
	report.Project.Branch = branch
	report.Project.Project = project
	report.Project.ApiVersion = apiversion
	report.NoOfCommits = len(commits)
	report.Commits = commits
	report.NoOfMergeRequests = len(mergeRequests)
	report.MergeRequests = mergeRequests
	report.NoOfMergeRequestCommits = len(mergeRequestCommits)
	report.MergeRequestCommits = mergeRequestCommits
	report.MergeRequestWithChanges = mergeRequestsWithChanges
	report.AmountOfMergeRequestCommits = (float64(report.NoOfMergeRequestCommits) / float64(report.NoOfCommits)) * float64(100)

	return report, nil
}
