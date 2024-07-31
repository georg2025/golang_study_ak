// this code will not work on golang 1.19 or earlier.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "your-access-token"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)
	gg := NewGeneralGithub(client)
	data, err := gg.GetItems(context.Background(), "ptflp", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
	data, err = gg.GetItems(context.Background(), "ptflp", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}

type GithubGist struct {
	GistList GistLister
}

type GithubRepo struct {
	RepoList RepoLister
}

type GeneralGithub struct {
	client *github.Client
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) (
		[]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist,
		*github.Response, error)
}

func NewGeneralGithub(client *github.Client) *GeneralGithub {
	g := &GeneralGithub{
		client: client,
	}

	return g
}

func NewGithubGist(client *github.Client) *GithubGist {
	g := &GithubGist{
		GistList: client.Gists,
	}

	return g
}

func NewGithubRepo(client *github.Client) *GithubRepo {
	g := &GithubRepo{
		RepoList: client.Repositories,
	}

	return g
}

func (gist *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := gist.GistList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	gsts := []Item{}
	for _, gist := range gists {
		gsts = append(gsts, Item{Title: *gist.ID, Description: *gist.Description, Link: *gist.GitPullURL})
	}

	return gsts, err
}

func (repo *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := repo.RepoList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	reps := []Item{}
	for _, rep := range repos {
		reps = append(reps, Item{Title: *rep.Name, Description: *rep.Description, Link: *rep.CloneURL})
	}

	return reps, err
}

func (general *GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

type Item struct {
	Title       string
	Description string
	Link        string
}
