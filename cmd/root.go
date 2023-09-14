package cmd

import (
	"fmt"
	"os"
	"prism/internal/service"

	"github.com/spf13/cobra"
)

var services *service.Service

var rootCmd = &cobra.Command{
	Use:   "prism",
	Short: "Creates a deployment template and configures the nomad",
	Long:  `Prism creates a deployment template and configures the nomad.`,
}

func Execute(service *service.Service) {
	services = service

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("failed to execute prism, %s", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prism.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
