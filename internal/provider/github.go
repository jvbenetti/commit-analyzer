package provider

import (
	"context"
	"setup/internal/analyzer"
	"setup/models"

	"github.com/google/go-github/v50/github"
)

func SearchCommits(user, repo string) (*models.Metric, error) {
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
		// O commit pode ser nil, é bom checar em Go
		if commit.Commit != nil && commit.Commit.Message != nil {
			msg := *commit.Commit.Message
			type_ := analyzer.CommitAn(msg) // Chama nossa lógica do passo 1
			result.Type[type_]++
		}
	}
	return &result, nil
}
