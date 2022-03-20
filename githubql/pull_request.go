package githubql

import (
	"context"
	"time"

	"github.com/shurcooL/githubv4"
)

type PullRequestFragment struct {
	ID        string
	Title     string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	Author    struct {
		Login string `json:"login"`
	}
	Reviews  ReviewNodes  `graphql:"reviews(first: 10)"`
	Comments CommentNodes `graphql:"comments(first: 10)"`
}

type ReviewNodes struct {
	Nodes []Review
}

type Review struct {
	Id     string
	Body   string
	State  string
	Url    string
	Author struct {
		Login string `json:"login"`
	}
	CreatedAt time.Time
}

type CommentNodes struct {
	Nodes []Comment
}

type Comment struct {
	Id     string
	Body   string
	Author struct {
		Login string `json:"login"`
	}
	CreatedAt time.Time
}

func (s Service) GetSalesEnginePullRequests(first int, after string) ([]PullRequestFragment, error) {
	var pullRequests []PullRequestFragment

	queryString := "repo:sephora-asia/luxola is:open author:carlqt author:Andrew2ww author:jawt94 author:ivan-fadieiev author:xinyingchua author:huilinn93"
	searchQuery, err := s.Search(queryString, first, after)

	for _, node := range searchQuery.Search.Nodes {
		pullRequests = append(pullRequests, node.PullRequest)
	}

	return pullRequests, err
}

// TODO: Make this return a PullRequest instead
func GetSalesEnginePullRequests(client *githubv4.Client) (SearchQuery, error) {
	var query SearchQuery

	variables := map[string]interface{}{
		"q":     githubv4.String("repo:sephora-asia/luxola is:open author:carlqt author:Andrew2ww author:jawt94 author:ivan-fadieiev author:xinyingchua author:huilinn93"),
		"first": githubv4.Int(10),
		"type":  githubv4.SearchType("ISSUE"),
	}

	err := client.Query(context.Background(), &query, variables)

	return query, err
}
