package models

import (
	"time"

	"github.com/carlqt/github_dashboard/githubql"
	"github.com/shurcooL/githubv4"
)

type PullRequests []PullRequest

type PullRequest struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    Author    `json:"author"`
	Reviews   []Review  `json:"reviews,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
}

type PullRequestModel struct {
	Client *githubv4.Client
}

func (pr PullRequestModel) All(perPage int, after string) (PullRequests, error) {
	// query, err := githubql.GetSalesEnginePullRequests(pr.Client)
	ghService := githubql.Service(pr)
	pullRequestFragments, err := ghService.GetSalesEnginePullRequests(perPage, after)

	if err != nil {
		return nil, err
	}

	return getPullRequests(pullRequestFragments), err
}

func (pr PullRequestModel) AllFromLogin(login string, perPage int, after string) (PullRequests, error) {
	ghService := githubql.Service(pr)
	pullRequestFragments, err := ghService.GetUserPullRequests(login, perPage, after)

	if err != nil {
		return nil, err
	}

	return getPullRequests(pullRequestFragments), err
}

func newPullRequest(pr githubql.PullRequestFragment) PullRequest {
	return PullRequest{
		ID:        pr.ID,
		Title:     pr.Title,
		Url:       pr.Url,
		Author:    pr.Author,
		CreatedAt: localizeTime(pr.CreatedAt),
		UpdatedAt: localizeTime(pr.UpdatedAt),
		Reviews:   newReviews(pr.Reviews.Nodes),
		Comments:  newComments(pr.Comments),
	}
}

func getPullRequests(fragments []githubql.PullRequestFragment) PullRequests {
	var pullRequests PullRequests

	for _, fragment := range fragments {
		pr := newPullRequest(fragment)
		pullRequests = append(pullRequests, pr)
	}

	return pullRequests
}

func localizeTime(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Singapore")

	return t.In(loc)
}
