package githubql

import (
	"context"

	"github.com/shurcooL/githubv4"
)

type SearchQuery struct {
	Search Search `graphql:"search(query: $q, type: $type, first: $first)"`
}

type Search struct {
	Nodes []struct {
		PullRequest PullRequestFragment `graphql:"... on PullRequest"`
	}
}

type Service struct {
	Client *githubv4.Client
}

func (s Service) Search(queryString string, first int, after string) (SearchQuery, error) {
	var query SearchQuery

	variables := map[string]interface{}{
		"q":     githubv4.String(queryString),
		"first": githubv4.Int(first),
		"type":  githubv4.SearchType("ISSUE"),
	}

	err := s.Client.Query(context.Background(), &query, variables)

	return query, err
}
