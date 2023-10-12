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
	Short: "Creating a nomad job configuration template.",
	Long:  `Prism creates a nomad job configuration template and deploys it to the cluster.`,
}

func Execute(service *service.Service) {
	services = service

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("error execute: %s", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
