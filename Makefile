.PHONY: build
build:
	go build -o releaser releaser.go


release:
	goreleaser