// Package github provides a GO API for the GitHub issue tracker.
package github

import (
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func (result *IssuesSearchResult) Len() int {
	return len(result.Items)
}

func (result *IssuesSearchResult) Swap(i, j int) {
	result.Items[i], result.Items[j] = result.Items[j], result.Items[i]
}

func (result *IssuesSearchResult) Less(i, j int) bool {
	return result.Items[i].CreatedAt.UnixNano() < result.Items[j].CreatedAt.UnixNano()
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
