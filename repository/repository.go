package repository

type Interface interface {
	GetRepositoryURL() string
	GetMergedBranches() []string
	GetCurrentBranch() string
	GetOwner() string
	GetName() string
}
