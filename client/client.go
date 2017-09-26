package client

type Interface interface {
	CreateRelease(version string) error
}
