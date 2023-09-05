package apitest

import (
	"fmt"
	"github.com/google/go-github/v54/github"
	"golang.org/x/net/context"
	"os"
	"testing"
)

// Fetch all issues
func fetchIssues(owner string, repo string) ([]*github.Issue, error) {
	client := github.NewClient(nil)
	issues, _, err := client.Issues.ListByRepo(context.Background(), owner, repo, nil)
	return issues, err
}

// Fetch issue all comment
func fetchIssueComments(owner string, repo string, number int) ([]*github.IssueComment, error) {
	client := github.NewClient(nil)
	comments, _, err := client.Issues.ListComments(context.Background(), owner, repo, number, nil)
	return comments, err
}

// Create test issue
func createIssue(token string, owner string, repo string, title string, body string) (*github.Issue, error) {
	client := github.NewTokenClient(context.Background(), token)
	request := &github.IssueRequest{Title: &title, Body: &body}
	issue, _, err := client.Issues.Create(context.Background(), owner, repo, request)
	return issue, err
}

// Add issue comment
func addIssueComment(token string, owner string, repo string, number int, body string) (*github.IssueComment, error) {
	client := github.NewTokenClient(context.Background(), token)
	comment := &github.IssueComment{
		Body: &body,
	}
	issueComment, _, err := client.Issues.CreateComment(context.Background(), owner, repo, number, comment)
	return issueComment, err
}

func TestIssues(t *testing.T) {
	t.Run("get issues", func(t *testing.T) {
		issues, err := fetchIssues("li-guohao", "test")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for i, issue := range issues {
			fmt.Printf("%v. %v\n", i+1, issue.GetTitle())
		}
	})
	t.Run("get issue comments", func(t *testing.T) {
		issuesComments, err := fetchIssueComments("li-guohao", "test", 3)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for i, comment := range issuesComments {
			fmt.Printf("%v. %v\n", i+1, comment.GetBody())
		}
	})
	t.Run("create issue", func(t *testing.T) {
		token := os.Getenv("GITHUB_TOKEN")
		issue, err := createIssue(token, "li-guohao", "test", "[Test] Test Title", "This is test issue \nbody create by github api.")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("number: %v - title: %v\n", issue.GetNumber(), issue.GetTitle())
	})
	t.Run("add issue comments", func(t *testing.T) {
		token := os.Getenv("GITHUB_TOKEN")
		comment, err := addIssueComment(token, "li-guohao", "test", 3, "This is new comment by api added.")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("%v. %v\n", comment.GetID(), comment.GetBody())
	})
}
