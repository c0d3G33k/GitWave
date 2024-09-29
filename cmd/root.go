package cmd

import (
	"os"

	"github.com/c0d3G33k/GitWave/pkg/common"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GitWave",
	Short: "Effortlessly clone, sync, and ride the latest changes across all your repositories",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&common.InputFilePath, "input", "i", "", "input file path")
	rootCmd.PersistentFlags().StringVarP(&common.WorkDirectory, "directory", "d", "/tmp/", "working directory path")

	rootCmd.MarkPersistentFlagRequired("input")
}
