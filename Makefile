.PHONY: build
build:
	go build -o releaser releaser.go

.PHONY: release
release: goreleaser glide
	goreleaser --rm-dist

.PHONY: goreleaser
goreleaser:
	@go get github.com/goreleaser/goreleaser && go install github.com/goreleaser/goreleaser

.PHONY: glide
glide:
	@go get github.com/Masterminds/glide && glide install