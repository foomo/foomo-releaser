package git

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var originalExecute func(dir, name string, args ...string) ([]string, error)

func setup() {
	originalExecute = execute
}

func teardown() {
	execute = originalExecute
}

func TestGit_GetRepositoryURL(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]string, error) {
		assert.Equal(t, "git", name)
		assert.Equal(t, "remote", args[0])
		assert.Equal(t, "-v", args[1])
		assert.Equal(t, "/www/globus", dir)
		return []string{
			"origin	git@github.com:bestbytes/globus.git (fetch)",
			"origin	git@github.com:bestbytes/globus.git (push)",
		}, nil
	}

	repo, _ := NewRepository("/www/globus")
	url := repo.GetRepositoryURL()
	assert.Equal(t, "https://github.com/bestbytes/globus", url)
}

func TestGit_GetMergedBranches(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]string, error) {
		assert.Equal(t, "git", name)
		assert.Equal(t, "branch", args[0])
		assert.Equal(t, "-r", args[1])
		assert.Equal(t, "--merged", args[2])
		assert.Equal(t, "/www/globus", dir)
		return []string{
			"origin/develop",
			"origin/feature/ECOMDEV-1-invalidate",
			"origin/feature/ECOMDEV-2-storefinder",
			"origin/feature/ECOMDEV-3-manual",
			"origin/release/1.20.0",
		}, nil
	}

	repo, _ := NewRepository("/www/globus")
	branches := repo.GetMergedBranches()

	assert.Equal(t, 3, len(branches))
	assert.Equal(t, "feature/ECOMDEV-1-invalidate", branches[0])
	assert.Equal(t, "feature/ECOMDEV-2-storefinder", branches[1])
	assert.Equal(t, "feature/ECOMDEV-3-manual", branches[2])
}

func TestGit_GetOwner(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]string, error) {
		return []string{
			"origin	git@github.com:bestbytes/globus.git (fetch)",
			"origin	git@github.com:bestbytes/globus.git (push)",
		}, nil
	}

	repo, _ := NewRepository("/www/globus")
	owner := repo.GetOwner()

	assert.Equal(t, "bestbytes", owner)
}

func TestGit_GetName(t *testing.T) {
	setup()
	defer teardown()

	execute = func(dir, name string, args ...string) ([]string, error) {
		return []string{
			"origin	git@github.com:bestbytes/globus.git (fetch)",
			"origin	git@github.com:bestbytes/globus.git (push)",
		}, nil
	}

	repo, _ := NewRepository("/www/globus")
	name := repo.GetName()

	assert.Equal(t, "globus", name)
}
