package client

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/smartinov/globus-release/repository"
	"fmt"
	"golang.org/x/net/context"
	"strings"
)

type githubClient struct {
	client     *github.Client
	ctx        context.Context
	repository repository.Interface
}

type githubContext struct {
}

var _ Interface = githubClient{}

func New(token string, repo repository.Interface) (Interface, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	ctx := context.Background()
	return githubClient{
		repository: repo,
		ctx:        ctx,
		client:     github.NewClient(oauth2.NewClient(context.Background(), ts)),
	}, nil
}

func (c githubClient) CreateRelease(version string) error {
	body := strings.Join(c.repository.GetMergedBranches(), "\n")

	var data = &github.RepositoryRelease{
		Name:            github.String("Release " + version),
		TagName:         github.String(version),
		TargetCommitish: github.String("release/" + version),
		Body:            github.String(body),
		Draft:           github.Bool(true),
	}

	_, _, err := c.client.Repositories.CreateRelease(c.ctx, c.repository.GetOwner(), c.repository.GetName(), data)
	if err != nil {
		return err
	}
	fmt.Print(data)
	return nil
}
