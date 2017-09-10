package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/zddhub/dit/dit"
)

var objectType, objectSize, objectContent bool

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file (-p | -t | -s) <object>",
	Short: "Provide content or type and size information for repository objects",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			return
		}

		repo := LoadRepository()
		object, content, err := repo.CatFile(args[0])

		if err != nil {
			fmt.Println("fatal: Not a valid object name", args[0])
			return
		}

		if objectType {
			fmt.Println(object.Type)
		}

		if objectSize {
			fmt.Println(object.Size)
		}

		if objectContent {
			fmt.Printf("%s", content)
		}
	},
}

func init() {
	ditCmd.AddCommand(catFileCmd)

	catFileCmd.Flags().BoolVarP(&objectSize, "size", "s", false, "show object size")
	catFileCmd.Flags().BoolVarP(&objectType, "type", "t", false, "show object type")
	catFileCmd.Flags().BoolVarP(&objectContent, "print", "p", false, "pretty-print object's content")
}
