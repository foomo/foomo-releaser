package git

import (
	"os/exec"
	"log"
	"strings"
	"fmt"
	"regexp"
	"github.com/smartinov/globus-release/repository"
)

const (
	githubTemplate           = "https://github.com/%s/%s"
	remoteRepositoryTemplate = "git@github.com:(\\w+)/(\\w+).git"
	mergedBranchesPrefix     = "feature/"
)

type git struct {
	dir string
}

var _ repository.Interface = git{}

func NewRepository(directory string) (repository.Interface, error) {
	return git{dir: directory}, nil
}

func (r git) GetRepositoryURL() string {
	return fmt.Sprintf(githubTemplate, r.GetOwner(), r.GetName())
}

func (r git) GetOwner() string {
	re := regexp.MustCompile(remoteRepositoryTemplate)
	return re.FindStringSubmatch(r.getRemote())[1]
}

func (r git) GetName() string {
	re := regexp.MustCompile(remoteRepositoryTemplate)
	return re.FindStringSubmatch(r.getRemote())[2]
}

func (r git) GetMergedBranches() []string {
	output, err := r.executeGit("branch", "-r", "--merged")
	if err != nil {
		log.Fatal(err)
	}
	var branches []string
	for _, line := range output {
		branch := strings.TrimPrefix(line, "origin/")
		if strings.HasPrefix(branch, mergedBranchesPrefix) {
			branches = append(branches, branch)
		}
	}
	return branches
}

func (r git) getRemote() string {
	output, err := r.executeGit("remote", "-v")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range output {
		if strings.HasPrefix(line, "origin") {
			return line
		}
	}
	return ""
}

func (r git) executeGit(arguments ... string) ([]string, error) {
	return execute(r.dir, "git", arguments...)
}

var execute = func(dir, commandName string, command ... string) ([]string, error) {
	cmd := exec.Command(commandName, command...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(output), "\n")
	for i := range lines {
		lines[i] = strings.Trim(lines[i], " ")
	}

	return lines, err
}
