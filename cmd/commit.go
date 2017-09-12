package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/zddhub/dit/dit"
)

var message string

// commitCmd represents the catFile command
var commitCmd = &cobra.Command{
	Use:   "commit -m <message>",
	Short: "Record changes to the repository",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(message) == 0 {
			fmt.Println("Aborting commit due to empty commit message.")
			return
		}
		repo := LoadRepository()
		repo.Commit(message)
	},
}

func init() {
	ditCmd.AddCommand(commitCmd)

	commitCmd.Flags().StringVarP(&message, "message", "m", "",
		"Use the given <message> as the commit message.")
}
