package client

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/foomo/foomo-releaser/repository"
	"golang.org/x/net/context"
	"fmt"
	"regexp"
	"strings"
	"log"
)

const (
	featureTicketRegex = "ECOMDEV-\\d+"
	featureTicketLink  = "https://jira.globuswiki.com/browse/%s"
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
	branch := strings.Split(c.repository.GetCurrentBranch(), "/")

	if !(len(branch) == 2 && branch[0] == "release" && branch[1] == version) {
		msg := "the specified repository needs to be set on release/%s for the scirpt to work, current branch '%s'"
		log.Fatalf(msg, version,c.repository.GetCurrentBranch())
	}

	var data = &github.RepositoryRelease{
		Name:            github.String("Release " + version),
		TagName:         github.String(version),
		TargetCommitish: github.String("release/" + version),
		Body:            github.String(c.getBody()),
		Draft:           github.Bool(true),
	}

	_, _, err := c.client.Repositories.CreateRelease(c.ctx, c.repository.GetOwner(), c.repository.GetName(), data)
	if err != nil {
		return err
	}
	return nil
}

func (c githubClient) getBody() string {
	body := "\n## Branches "
	for _, branch := range c.repository.GetMergedBranches() {
		body += fmt.Sprintf("\n - [%s](%s) ", branch, calculateJiraLink(branch))
	}
	return body
}

func calculateJiraLink(branch string) string {
	re := regexp.MustCompile(featureTicketRegex)
	ticket := re.FindString(branch)
	if ticket == "" {
		return ""
	}
	return fmt.Sprintf(featureTicketLink, ticket)
}
