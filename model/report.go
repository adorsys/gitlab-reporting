package model

type Report struct {
	Project                     Project
	NoOfCommits                 int
	Commits                     []Commit
	NoOfMergeRequests           int
	MergeRequests               []MergeRequest
	NoOfMergeRequestCommits     int
	MergeRequestCommits         []Commit
	AmountOfMergeRequestCommits float64
	NoOfChanges                 int
	MergeRequestWithChanges     []MergeRequestWithChanges
}
