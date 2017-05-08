package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/adorsys/gitlab-reporting/model"
)

var perPage = 100

func GetAllCommits(pURL, pToken, pBranch, pProj, pVApi *string, page int) ([]model.Commit, error) {

	strProj := strings.Replace(*pProj, "/", "%2F", -1)
	var url = fmt.Sprint(*pURL, "/api/v", *pVApi, "/projects/", strProj, "/repository/commits")

	url += fmt.Sprint("?per_page=", perPage, "&page=", page)
	if *pBranch != "" {
		url += fmt.Sprint("&sha=", *pBranch)
	}

	fmt.Printf("\nGetCommits(): from %v %v %v", url, *pToken, "...\n")

	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", *pToken)
	req.Close = false

	res, err := tr.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	commits := make([]model.Commit, 0)

	if err := json.Unmarshal(body, &commits); err != nil {
		return nil, err
	}

	if len(commits) == perPage {
		newCommits, err := GetAllCommits(pURL, pToken, pBranch, pProj, pVApi, page+1)

		if err != nil {
			return nil, err
		}

		for _, commit := range newCommits {
			commits = append(commits, commit)
		}
	}

	return commits, nil
}

func GetMergeRequests(pURL, pToken, pBranch, pProj, pVApi *string, page int) ([]model.MergeRequest, error) {

	strProj := strings.Replace(*pProj, "/", "%2F", -1)
	var url = fmt.Sprint(*pURL, "/api/v", *pVApi, "/projects/", strProj, "/merge_requests")

	// per_page max is 100
	url += fmt.Sprint("?per_page=", perPage, "&page=", page)
	if *pBranch != "" {
		url += fmt.Sprint("&sha=", *pBranch)
	}

	fmt.Printf("\nGetMergeRequests(): from %v %v %v", url, *pToken, "...\n")

	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", *pToken)
	req.Close = false

	res, err := tr.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	mergeRequests := make([]model.MergeRequest, 0)

	if err := json.Unmarshal(body, &mergeRequests); err != nil {
		return nil, err
	}

	if len(mergeRequests) == perPage {
		newMergeRequests, err := GetMergeRequests(pURL, pToken, pBranch, pProj, pVApi, page+1)

		if err != nil {
			return nil, err
		}

		for _, request := range newMergeRequests {
			mergeRequests = append(mergeRequests, request)
		}
	}

	return mergeRequests, nil
}

func GetMergeRequestCommits(pURL, pToken, pBranch, pProj, pVApi *string, mergeReqID int, page int) ([]model.Commit, error) {

	strProj := strings.Replace(*pProj, "/", "%2F", -1)
	var url = fmt.Sprint(*pURL, "/api/v", *pVApi, "/projects/", strProj, "/merge_requests/", mergeReqID, "/commits")

	// per_page max is 100
	url += fmt.Sprint("?per_page=", perPage, "&page=", page)
	if *pBranch != "" {
		url += fmt.Sprint("&sha=", *pBranch)
	}

	fmt.Printf("\nGetMergeRequestCommits(): from %v %v %v", url, *pToken, "...\n")

	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", *pToken)
	req.Close = false

	res, err := tr.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	mergeRequestCommits := make([]model.Commit, 0)

	if err := json.Unmarshal(body, &mergeRequestCommits); err != nil {
		return nil, err
	}

	if len(mergeRequestCommits) == perPage {
		newMergeRequestCommits, err := GetMergeRequestCommits(pURL, pToken, pBranch, pProj, pVApi, mergeReqID, page+1)

		if err != nil {
			return nil, err
		}

		for _, commit := range newMergeRequestCommits {
			mergeRequestCommits = append(mergeRequestCommits, commit)
		}
	}
	return mergeRequestCommits, nil
}

func GetMergeRequestChanges(pURL, pToken, pBranch, pProj, pVApi *string, mergeReqID int, page int) (model.MergeRequestWithChanges, error) {

	var mergeRequestWithChanges model.MergeRequestWithChanges
	strProj := strings.Replace(*pProj, "/", "%2F", -1)
	var url = fmt.Sprint(*pURL, "/api/v", *pVApi, "/projects/", strProj, "/merge_requests/", mergeReqID, "/changes")

	// per_page max is 100
	url += fmt.Sprint("?per_page=", perPage, "&page=", page)
	if *pBranch != "" {
		url += fmt.Sprint("&sha=", *pBranch)
	}

	fmt.Printf("\nGetMergeRequestChanges(): from %v %v %v", url, *pToken, "...\n")

	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", *pToken)
	req.Close = false

	res, err := tr.RoundTrip(req)

	if err != nil {
		return mergeRequestWithChanges, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &mergeRequestWithChanges); err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			log.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *json.InvalidUnmarshalError:
			log.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			log.Println(err)
		}
		return mergeRequestWithChanges, err

	}
	return mergeRequestWithChanges, nil
}
