package cmd

import (
	"errors"
	"os"
	"path"
	"strings"

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
	if len(remotes) == 0 {
		return "", errors.New("no remotes there")
	}
	remoteConfig := remotes[0].Config()
	if len(remoteConfig.URLs) == 0 {
		return "", errors.New("no URLs in remote-config")
	}
	url := remoteConfig.URLs[0]
	url = strings.TrimSuffix(url, ".git")
	return url, nil
}

func CreateReadme() utility.Readme {
	title := utility.RequestValueInput("title", nil)
	description := utility.RequestValueInput("description", nil)

	predictedRepoUrl, err := getRepoUrl()
	var repoURL string
	if err == nil {
		repoURL = utility.RequestValueInput("repository URL", predictedRepoUrl)
	} else {
		repoURL = utility.RequestValueInput("repository URL", nil)
	}
	repoURL = strings.TrimSuffix(repoURL, "/")

	watermark := utility.RequestDecisionInput("show watermark", true)

	return utility.Readme{
		Title:         title,
		Description:   description,
		RepositoryURL: repoURL,
		Watermark:     watermark,
	}
}
