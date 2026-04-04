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
	// Slice to get list of commits
	var allCommits []*github.RepositoryCommit

	// For loop, page -> page
	for {
		commits, resp, err := client.Repositories.ListCommits(ctx, user, repo, opt)
		if err != nil {
			return nil, err
		}

		// Get new commits and append in list
		allCommits = append(allCommits, commits...)

		if resp.NextPage == 0 {
			break
		}

		// Go to next "page"
		opt.Page = resp.NextPage
	}

	result := models.Metric{ // Call Metric Model
		Total: len(allCommits),
		Type:  make(map[string]int),
	}
	for _, commit := range allCommits {
		// Check is commit is nil, good practice
		if commit.Commit != nil && commit.Commit.Message != nil {
			msg := *commit.Commit.Message
			type_ := analyzer.CommitAn(msg) // Call analyze logic func
			result.Type[type_]++
		}
	}
	return &result, nil // Return a pointer and could be nil
}

func SearchAllCommitsChanges(user, repo string) ([]models.CommitChanges, error) {
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	var client *github.Client

	if token == "" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}

	// With pagination to get ALL commits
	opt := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var allCommits []*github.RepositoryCommit

	for {
		commits, resp, err := client.Repositories.ListCommits(ctx, user, repo, opt)
		if err != nil {
			return nil, err
		}

		allCommits = append(allCommits, commits...)

		if resp.NextPage == 0 || len(allCommits) >= 200 {
			break
		}

		// Go to the next page
		opt.Page = resp.NextPage
	}

	// Max 200 items
	if len(allCommits) > 200 {
		allCommits = allCommits[:200]
	}

	var results []models.CommitChanges

	// Iter about 200 commit to get changes
	for _, commit := range allCommits {
		if commit.SHA == nil {
			continue
		}

		singleCommit, _, err := client.Repositories.GetCommit(ctx, user, repo, *commit.SHA, nil)
		if err != nil {
			continue
		}

		changes := 0
		if singleCommit.Stats != nil {
			changes = singleCommit.Stats.GetTotal()
		}

		results = append(results, models.CommitChanges{
			SHA:     *commit.SHA,
			Changes: changes,
		})
	}

	return results, nil
}
