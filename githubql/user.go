package githubql

import (
	"context"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
)

type UserQuery struct {
	User UserFragment `graphql:"user(login: $login)"`
}

type UserFragment struct {
	Id                      string
	Name                    string
	Login                   string
	ContributionsCollection ContributionsCollection `graphql:"contributionsCollection(organizationID: $orgID)"`
}

type ContributionsCollection struct {
	TotalPullRequestReviewContributions int
	PullRequestReviewContributions      PullRequestReviewContributions `graphql:"pullRequestReviewContributions(first: 10)"`
}

type PullRequestReviewContributions struct {
	Edges []PullRequestEdge
}

type PullRequestEdge struct {
	Node PullRequestNode
}

type PullRequestNode struct {
	OccurredAt  time.Time
	PullRequest struct {
		Title          string
		BaseRepository struct {
			Name string
		}
	}
	PullRequestReview Review
}

func GetUser(client *githubv4.Client, login string) (UserFragment, error) {
	var query UserQuery
	orgID := os.Getenv("ORGANIZATION_ID")

	variables := map[string]interface{}{
		"login": githubv4.String(login),
		"orgID": githubv4.ID(orgID),
	}

	err := client.Query(context.Background(), &query, variables)
	return query.User, err
}
