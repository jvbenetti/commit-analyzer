package provider

import (
	"context"
	"os"
	"setup/internal/analyzer"
	"setup/models"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

// SearchCommits func return a pointer for optimization
func SearchCommits(user, repo string) (
	*models.Metric,
	error,
) {
	token := os.Getenv("GITHUB_TOKEN")

	ctx := context.Background()
	var client *github.Client

	if token == "" {
		// Local token
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client = github.NewClient(tc)
	}
	opt := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	commits, _, err := client.Repositories.ListCommits(ctx, user, repo, nil)
	if err != nil {
		return nil, err
	}

	result := models.Metric{ // Call Metric Model
		Total: len(commits),
		Type:  make(map[string]int),
	}
	for _, commit := range commits {
		// Check is commit is nil, good practice
		if commit.Commit != nil && commit.Commit.Message != nil {
			msg := *commit.Commit.Message
			type_ := analyzer.CommitAn(msg) // Call analyze logic func
			result.Type[type_]++
		}
	}
	return &result, nil // Return a pointer and could be nil
}
