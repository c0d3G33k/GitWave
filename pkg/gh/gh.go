package gh

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/c0d3G33k/GitWave/pkg/common"
	"github.com/c0d3G33k/GitWave/pkg/common/ui"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GetRepoNameFromURL(repoURL string) (name string, err error) {
	parsedURL, err := url.Parse(repoURL)

	if err != nil {
		return "", err
	}

	splitURL := strings.Split(parsedURL.Path, "/")

	if len(splitURL) == 0 {
		return "", fmt.Errorf("invalid URL")
	}

	return splitURL[len(splitURL)-1], nil
}

func SyncRepo(repoURL string) {
	repoName, errRepoName := GetRepoNameFromURL(repoURL)

	if errRepoName != nil {
		ui.PrintError("Error parsing repository URL, Skipping...")
	}

	directory := common.WorkDirectory + repoName

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		_, errClone := git.PlainClone(directory, false, &git.CloneOptions{
			URL:          repoURL,
			Progress:     nil,
			SingleBranch: true,
			Auth: &http.BasicAuth{
				Username: "username",
				Password: os.Getenv("GITHUB_PAT"),
			},
		})

		if errClone != nil {
			ui.PrintError("Error cloning repository : %v, %s", errClone, repoURL)
		} else {
			ui.PrintSuccess("Successfully cloned repository : %s", repoURL)
		}

	} else {
		repo, errOpen := git.PlainOpen(directory)

		if errOpen != nil {
			ui.PrintError("Error opening repository : %v, %s", errOpen, repoURL)
		}

		workTree, errWorkTree := repo.Worktree()

		if errWorkTree != nil {
			ui.PrintError("Error getting worktree : %v, %s", errWorkTree, repoURL)
		}

		errPull := workTree.Pull(&git.PullOptions{
			RemoteName: "origin",
			Auth: &http.BasicAuth{
				Username: "username",
				Password: os.Getenv("GITHUB_PAT"),
			},
		})

		if errPull != nil && errPull != git.NoErrAlreadyUpToDate {
			ui.PrintError(fmt.Sprintf("Error pulling repository : %v, %s", errPull, repoURL))
		} else if errPull == git.NoErrAlreadyUpToDate {
			ui.PrintWarning("Repository is already up to date : %s", repoURL)
		} else {
			ui.PrintSuccess("Successfully pulled repository : %s", repoURL)
		}
	}
}
