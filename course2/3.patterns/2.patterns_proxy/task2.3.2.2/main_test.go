package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v53/github"
)

type MockGistLister struct {
	GistsFunc func(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist,
		*github.Response, error)
}

func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist,
	*github.Response, error) {
	if m.GistsFunc != nil {
		return m.GistsFunc(ctx, username, nil)
	}

	return nil, nil, nil
}

type MockRepoLister struct {
	RepoFunc func(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository,
		*github.Response, error)
}

func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository,
	*github.Response, error) {
	if m.RepoFunc != nil {
		return m.RepoFunc(ctx, username, nil)
	}

	return nil, nil, nil
}

func TestRepoList(t *testing.T) {
	mockRepoService := &MockRepoLister{
		RepoFunc: func(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
			return []*github.Repository{
				{
					Name:        github.String("1"),
					Description: github.String("Test Repo 1"),
					CloneURL:    github.String("http:"),
				},
				{
					Name:        github.String("2"),
					Description: github.String("Test Repo 2"),
					CloneURL:    github.String("http1:"),
				},
				{
					Name:        github.String("3"),
					Description: github.String("Test Repo 3"),
					CloneURL:    github.String("http2:"),
				},
				{
					Name:        github.String("4"),
					Description: github.String("Test Repo 4"),
					CloneURL:    github.String("http3g:"),
				},
			}, nil, nil
		},
	}

	adapter := GithubAdapter{RepoList: mockRepoService}
	proxy := NewProxy(&adapter)
	ctx := context.Background()
	repos, err := proxy.GetRepos(ctx, "test")

	if err != nil {
		t.Errorf("expected no error, got error: %v", err)
	}

	if len(proxy.cache) != 1 {
		t.Errorf("wrong work with cache. Len should be 1, but it is: %d", len(proxy.cache))
	}

	if len(repos) != 4 {
		t.Errorf("expected 2 gists, got %d", len(repos))
	}

	if repos[0].Title != "1" || repos[1].Description != "Test Repo 2" {
		t.Errorf("unexpected gists: %+v", repos)
	}
}

func TestGistList(t *testing.T) {
	mockGistService := &MockGistLister{
		GistsFunc: func(ctx context.Context, user string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
			return []*github.Gist{
				{
					ID:          github.String("1"),
					Description: github.String("Test Gist 1"),
					GitPullURL:  github.String("http:"),
				},
				{
					ID:          github.String("2"),
					Description: github.String("Test Gist 2"),
					GitPullURL:  github.String("http1:"),
				},
				{
					ID:          github.String("3"),
					Description: github.String("Test Gist 3"),
					GitPullURL:  github.String("http2:"),
				},
				{
					ID:          github.String("4"),
					Description: github.String("Test Gist 4"),
					GitPullURL:  github.String("http3g:"),
				},
			}, nil, nil
		},
	}

	adapter := GithubAdapter{GistList: mockGistService}
	proxy := NewProxy(&adapter)
	ctx := context.Background()
	gists, err := proxy.GetGists(ctx, "test")

	if err != nil {
		t.Errorf("expected no error, got error: %v", err)
	}

	if len(proxy.cache) != 1 {
		t.Errorf("wrong work with cache. Len should be 1, but it is: %d", len(proxy.cache))
	}

	if len(gists) != 4 {
		t.Errorf("expected 2 gists, got %d", len(gists))
	}

	if gists[0].Description != "Test Gist 1" || gists[1].Description != "Test Gist 2" {
		t.Errorf("unexpected gists: %+v", gists)
	}
}

func TestGistRepoErrorOccur(t *testing.T) {
	mockRepoService := &MockRepoLister{
		RepoFunc: func(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
			return nil, nil, fmt.Errorf("error here")
		},
	}

	adapter := GithubAdapter{RepoList: mockRepoService}
	proxy := NewProxy(&adapter)
	ctx := context.Background()
	_, err := proxy.GetRepos(ctx, "test")

	if err == nil {
		t.Errorf("expected error, got none")
	}

	mockGistService := &MockGistLister{
		GistsFunc: func(ctx context.Context, user string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
			return nil, nil, fmt.Errorf("error here")
		},
	}

	adapter.GistList = mockGistService
	_, err = proxy.GetGists(ctx, "test")

	if err == nil {
		t.Errorf("expected error, got none")
	}
}
