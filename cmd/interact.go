package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/Tch1b0/readcli/pkg/utility"
	"github.com/go-git/go-git/v5"
)

func getRepoUrl() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	repo, err := git.PlainOpen(path.Join(dir, ".git"))
	if err != nil {
		return "", err
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return "", err
	}

	url := remotes[0].Config().URLs[0]
	return url, nil
}

func CreateReadme() utility.Readme {
	predictedRepoUrl, err := getRepoUrl()
	var repoUrl interface{}
	if err == nil {
		repoUrl = utility.RequestValueInput("repository URL", predictedRepoUrl)
	} else {
		repoUrl = utility.RequestValueInput("repository URL", nil)
	}
	fmt.Println(repoUrl)
	return utility.Readme{}
}
