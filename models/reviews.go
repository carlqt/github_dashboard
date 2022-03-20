package models

import (
	"time"

	"github.com/carlqt/github_dashboard/githubql"
)

type Review struct {
	Id        string    `json:"id,omitempty"`
	Body      string    `json:"body"`
	Url       string    `json:"url"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"createdAt"`
	Author    Author    `json:"author"`
}

func newReviewsFromGithubql(pr githubql.PullRequestReviewContributions) []Review {
	var reviews []Review

	for _, r := range pr.Edges {
		review := Review{
			Id:        r.Node.PullRequestReview.Id,
			Body:      r.Node.PullRequestReview.Body,
			State:     r.Node.PullRequestReview.State,
			Url:       r.Node.PullRequestReview.Url,
			CreatedAt: localizeTime(r.Node.PullRequestReview.CreatedAt),
		}
		reviews = append(reviews, review)
	}

	return reviews
}

func newReviews(ghReviews []githubql.Review) []Review {
	var reviews []Review

	for _, r := range ghReviews {
		review := Review{
			Body:      r.Body,
			State:     r.State,
			Url:       r.Url,
			Id:        r.Id,
			Author:    r.Author,
			CreatedAt: r.CreatedAt,
		}

		reviews = append(reviews, review)
	}

	return reviews
}
