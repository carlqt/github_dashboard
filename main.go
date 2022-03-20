package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/carlqt/github_dashboard/models"
	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"github.com/wailsapp/wails"
	"golang.org/x/oauth2"
)

func basic() string {
	return "Hello World!"
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}

func ghClient() *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	return githubv4.NewClient(httpClient)
}

func userPullRequests(login string) (models.PullRequests, error) {
	client := ghClient()
	model := models.PullRequestModel{Client: client}

	return model.AllFromLogin(login, 5, "")
}

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "github_dashboard",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(userPullRequests)
	app.Run()
}
