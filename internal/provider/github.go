package provider

import (
	"context"
	"setup/internal/analyzer"
	"setup/models"

	"github.com/google/go-github/v50/github"
)

// SearchCommits func return a pointer for optimization
func SearchCommits(user, repo string) (
	*models.Metric,
	error,
) {
	client := github.NewClient(nil)
	ctx := context.Background()

	commits, _, err := client.Repositories.ListCommits(ctx, user, repo, nil)
	if err != nil {
		return nil, err
	}

	result := models.Metric{
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
