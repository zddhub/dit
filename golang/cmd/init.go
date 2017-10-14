package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zddhub/dit/golang/dit"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty Dit repository or reinitialize an existing one",
	Run: func(cmd *cobra.Command, args []string) {
		repo := dit.NewRepository()
		repo.Init()
	},
}

func init() {
	ditCmd.AddCommand(initCmd)
}
