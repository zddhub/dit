package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// ditCmd represents the base command when called without any subcommands
var ditCmd = &cobra.Command{
	Use:   "dit",
	Short: "dit - the stupid content tracker like git",
	Long:  "usage: dit [--help] <command> [<args>]",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the ditCmd.
func Execute() {
	if err := ditCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const usageTemplate = `{{.Short}}

{{if .HasAvailableSubCommands}}These are common Dit commands used in various situations:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{if .HasAvailableLocalFlags}}Options: 
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{end}}

See 'dit help <command>' to read about a specific subcommand.
`

func init() {
	ditCmd.SetUsageTemplate(usageTemplate)
}
