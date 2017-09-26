package main

import (
	"flag"
	"github.com/smartinov/globus-releaser/client"
	"os"
	"path/filepath"
	"log"
	"github.com/smartinov/globus-releaser/repository/git"
	"fmt"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	token := flag.String("t", "", "github token to use when publishing repositories")
	directory := flag.String("d", dir, "directory of the repository to release (default: current directory)")
	version := flag.String("v", "", "version of the release to publish")

	flag.Parse()
	if *token == "" {
		log.Fatal("token parameter for github must be provided")
	}
	if !directoryExists(*directory) {
		log.Fatal("specified directory must exist and must be a git repository")
	}

	repo, err := git.NewRepository(*directory)
	if err != nil {
		log.Fatal("could not initialize repository with err ", err.Error())
	}

	ghc, err := client.New(*token, repo)
	if err != nil {
		log.Fatal("could not initialize the github client ")
	}

	err = ghc.CreateRelease(*version)
	if err != nil {
		log.Fatal("error occured while creating a release", err.Error())
	}

	fmt.Println("sucessfully created release " + *version)
}

// exists returns whether the given file or directory exists or not
func directoryExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
