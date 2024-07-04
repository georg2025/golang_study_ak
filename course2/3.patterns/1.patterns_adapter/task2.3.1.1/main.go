// important info: this code wont start on go 1.19, cause this version is not supported. At the same time
// it can be easily used on 20th version and so on.
package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "your_acess_token"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	g := NewGithubAdapter(client)
	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "ptflp"))
}

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) (
		[]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist,
		*github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
	// opt :=&github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
}

type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}
	return g
}

func (adapter *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := adapter.GistList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	gsts := []Item{}
	for _, gist := range gists {
		gsts = append(gsts, Item{Title: *gist.ID, Description: *gist.Description, Link: *gist.GitPullURL})
	}

	return gsts, err
}

func (adapter *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := adapter.RepoList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	reps := []Item{}
	for _, repo := range repos {
		reps = append(reps, Item{Title: *repo.Name, Description: *repo.Description, Link: *repo.CloneURL})
	}

	return reps, err
}

type Item struct {
	Title       string
	Description string
	Link        string
}
