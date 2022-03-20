package models

import (
	"time"

	"github.com/carlqt/github_dashboard/githubql"
)

type Comment struct {
	Id        string    `json:"id"`
	Body      string    `json:"body"`
	Author    Author    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

func newComments(commentNodes githubql.CommentNodes) []Comment {
	var comments []Comment

	for _, c := range commentNodes.Nodes {
		comment := Comment{
			Id:        c.Id,
			Body:      c.Body,
			Author:    c.Author,
			CreatedAt: localizeTime(c.CreatedAt),
		}

		comments = append(comments, comment)
	}

	return comments
}
