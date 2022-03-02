package cmd

import (
	"errors"
	"os"
	"path"
	"path/filepath"
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
	url = strings.TrimSuffix(url, "/")
	url = strings.TrimSuffix(url, ".git")
	return url, nil
}

func getCurrentDirectory() (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	_, tmp := filepath.Split(currentPath)
	return tmp, nil
}

func CreateReadme() (utility.Readme, error) {
	dir, err := getCurrentDirectory()
	if err != nil {
		return utility.Readme{}, err
	}

	predictedRepoUrl, err := getRepoUrl()
	var repoURL string
	if err == nil {
		repoURL = utility.RequestValueInput("repository URL", predictedRepoUrl)
	} else {
		repoURL = utility.RequestValueInput("repository URL", nil)
	}
	repoURL = strings.TrimSuffix(repoURL, "/")

	return utility.Readme{
		Title:            utility.RequestValueInput("title", dir),
		Description:      utility.RequestValueInput("description", nil),
		RepositoryURL:    repoURL,
		Requirements:     utility.RequestValueListInput("requirements"),
		ShowShields:      utility.RequestDecisionInput("show shields", true),
		ShowContributors: utility.RequestDecisionInput("list contributors", false),
		Watermark:        utility.RequestDecisionInput("show watermark", true),
	}, nil
}
