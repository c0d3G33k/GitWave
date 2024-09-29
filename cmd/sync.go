package cmd

import (
	"os"

	"github.com/c0d3G33k/GitWave/pkg/common/ui"
	"github.com/c0d3G33k/GitWave/pkg/common/utils"
	"github.com/c0d3G33k/GitWave/pkg/gh"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync list of repositories",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if !utils.CheckGitHubPAT() {
			ui.PrintError("Please set GITHUB_PAT environment variable")
			os.Exit(1)
		}

		gh.Connecter()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
