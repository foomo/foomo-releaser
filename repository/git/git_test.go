package git

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"errors"
)

const (
	repositoryDirectory = "/www/globus"
)

var originalExecute func(dir, name string, args ...string) ([]byte, error)

func setup() {
	originalExecute = execute
}

func teardown() {
	execute = originalExecute
}

func TestGit_GetRepositoryURL(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]byte, error) {
		assert.Equal(t, repositoryDirectory, dir)
		return []byte(strings.Join([]string{
			"origin	git@github.com:bestbytes/globus.git (fetch)",
			"origin	git@github.com:bestbytes/globus.git (push)",
		}, "\n")), nil
	}

	repo, _ := NewRepository(repositoryDirectory)
	url := repo.GetRepositoryURL()
	assert.Equal(t, "https://github.com/bestbytes/globus", url)
}

func TestGit_GetMergedBranches(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]byte, error) {
		assert.Equal(t, repositoryDirectory, dir)
		return []byte(strings.Join([]string{
			"origin/develop",
			"origin/feature/ECOMDEV-1-invalidate",
			"origin/feature/ECOMDEV-2-storefinder",
			"origin/feature/ECOMDEV-3-manual",
			"origin/release/1.20.0",
		}, "\n")), nil
	}

	repo, _ := NewRepository(repositoryDirectory)
	branches := repo.GetMergedBranches()

	assert.Equal(t, 3, len(branches))
	assert.Equal(t, "feature/ECOMDEV-1-invalidate", branches[0])
	assert.Equal(t, "feature/ECOMDEV-2-storefinder", branches[1])
	assert.Equal(t, "feature/ECOMDEV-3-manual", branches[2])
}

func TestGit_GetOwner(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]byte, error) {
		return []byte(strings.Join([]string{
			"origin	git@github.com:bestbytes/globus.git (fetch)",
			"origin	git@github.com:bestbytes/globus.git (push)",
		}, "\n")), nil
	}

	repo, _ := NewRepository(repositoryDirectory)
	owner := repo.GetOwner()

	assert.Equal(t, "bestbytes", owner)
}

func TestGit_GetName(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]byte, error) {
		return []byte(strings.Join([]string{
			"origin	git@github.com:bestbytes/globus.git (fetch)",
			"origin	git@github.com:bestbytes/globus.git (push)",
		}, "\n")), nil
	}

	repo, _ := NewRepository(repositoryDirectory)
	name := repo.GetName()

	assert.Equal(t, "globus", name)
}

func TestGit_GetCurrentBranch(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]byte, error) {
		assert.Equal(t, repositoryDirectory, dir)
		return []byte("release/123"), nil
	}

	repo, _ := NewRepository(repositoryDirectory)
	branch := repo.GetCurrentBranch()

	assert.Equal(t, "release/123", branch)
}

func TestNewRepositoryErr(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]byte, error) {
		return nil, errors.New("NOUP")
	}

	_, err := NewRepository(repositoryDirectory)
	assert.NotNil(t, err)

}
