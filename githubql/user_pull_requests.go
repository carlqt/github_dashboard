package githubql

import "fmt"

func (s Service) GetUserPullRequests(login string, first int, after string) ([]PullRequestFragment, error) {
	var pullRequests []PullRequestFragment

	queryString := fmt.Sprintf("repo:sephora-asia/luxola author:%s", login)
	searchQuery, err := s.Search(queryString, first, after)

	for _, node := range searchQuery.Search.Nodes {
		pullRequests = append(pullRequests, node.PullRequest)
	}

	return pullRequests, err
}
