package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create prism project",
	Long: fmt.Sprintf(
		"Create new deployment project.\n%s",
		"Creates a project directory with default configuration files.",
	),
	Run: func(cmd *cobra.Command, args []string) {
		// Project name.
		var name string

		if len(args) > 0 {
			name = args[0]
		}

		// Create a project.
		projectName, err := services.Project.Create(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Project \"%s\" successfully created.\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
