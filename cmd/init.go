package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project",
	Long: fmt.Sprintf(
		"Create new deployment project.\n%s\n%s %s",
		"Creates a project directory with default configuration files.",
		"To set your project name, specify it after the init command:",
		"prism init <project-name>",
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
			fmt.Printf("An error occurred while creating the project: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Project \"%s\" successfully created.\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
