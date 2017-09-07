package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/zddhub/dit/dit"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [<options>] [--] <pathspec>...",
	Short: "Add file contents to the index",
	Run: func(cmd *cobra.Command, args []string) {
		repo := NewRepository()
		repo.AddFiles(args)
	},
}

func init() {
	ditCmd.AddCommand(addCmd)
}
