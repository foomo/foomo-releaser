package repository

type Interface interface {
	GetRepositoryURL() string
	GetMergedBranches() []string
	GetOwner() string
	GetName() string
}
