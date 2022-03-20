package models

import (
	"github.com/carlqt/github_dashboard/githubql"
	"github.com/shurcooL/githubv4"
)

type Author struct {
	Login string `json:"login"`
}

type User struct {
	ID      string   `json:"id"`
	Login   string   `json:"login"`
	Name    string   `json:"name"`
	Reviews []Review `json:"reviews"`
}

type Users []User
type UserModel struct {
	Client *githubv4.Client
}

func (u UserModel) GetUserByLogin(login string) (User, error) {
	query, err := githubql.GetUser(u.Client, login)

	if err != nil {
		return User{}, err
	}

	return newUser(query), err
}

func newUser(u githubql.UserFragment) User {
	return User{
		ID:      u.Id,
		Name:    u.Name,
		Login:   u.Login,
		Reviews: newReviewsFromGithubql(u.ContributionsCollection.PullRequestReviewContributions),
	}
}
