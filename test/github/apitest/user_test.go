package apitest

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v54/github"
)

// Fetch all the public organizations' membership of a user.
func fetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func TestUser(t *testing.T) {
	t.Run("get user organizations", func(t *testing.T) {
		username := "li-guohao"
		organizations, err := fetchOrganizations(username)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for i, organization := range organizations {
			fmt.Printf("%v. %v - Desc: %v - Followers: %v\n", i+1, organization.GetLogin(), organization.GetDescription(), organization.GetFollowers())
		}
	})
}
