package gh

import (
	"os"

	"github.com/c0d3G33k/GitWave/pkg/common"
	"github.com/c0d3G33k/GitWave/pkg/common/ui"
	"github.com/c0d3G33k/GitWave/pkg/common/utils"
)

func Connecter() {
	utils.ValidateWorkDirectory()
	repoList, errReadFile := utils.ReadCSVFile(common.InputFilePath)

	if errReadFile != nil {
		ui.PrintError("Error reading input file : %v", errReadFile)
		os.Exit(1)
	}

	for _, repo := range repoList {
		repoURL := repo[0]

		SyncRepo(repoURL)
	}
}
