package cmd

import (
	"fmt"
	"prism/internal/service"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "out",
	Short: "Create nomad configuration file",
	Long:  `Create nomad template configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		path, _ := cmd.Flags().GetString("path")
		from, _ := cmd.Flags().GetString("from")

		if name != "" && path != "" {
			var service service.OutputService
			service.CreateNomadConfiguration(name, path, from)
			return
		}

		fmt.Println("file name and path must be specified")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name output file")
	createCmd.Flags().StringP(
		"path",
		"p",
		"",
		"directory where the configuration \"nomad\" file will be created",
	)
	createCmd.Flags().StringP(
		"from",
		"f",
		"",
		"path to configuration \"yaml\" file",
	)
}
