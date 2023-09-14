package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"prism/config"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create prism project",
	Long: fmt.Sprintf(
		"Create new deployment project.\n%s.",
		"Creates a project directory with default configuration files.",
	),
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags.
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("failed to read flag \"name\", %s\n", err)
			os.Exit(1)
		}

		// Get root dir path.
		rootDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("failed project initialization, %s\n", err)
			os.Exit(1)
		}

		// Create project directories.
		projectDir := "prism"
		if name != "" {
			projectDir = name
		}

		projectDirPath := filepath.Join(rootDir, projectDir)

		projectFileDir := "files"
		fileDirPath := filepath.Join(projectDirPath, projectFileDir)

		chartName := strings.ReplaceAll(projectDir, "-", "_")
		chartFileName := fmt.Sprintf("%s.yaml", chartName)

		configName := "config"
		configFileName := fmt.Sprintf("%s.yaml", configName)

		dirStat, err := os.Stat(projectDirPath)
		if err != nil || !dirStat.IsDir() {
			err = os.MkdirAll(
				filepath.Join(projectDir, projectFileDir),
				0700,
			)

			if err != nil {
				fmt.Printf("failed project initialization, %s\n", err)
				os.Exit(1)
			}
		}

		// Create default project files.
		// chart.
		err = services.Project.CreateDefautlFile(
			config.ConfigFile,
			"prism.yaml",
			chartFileName,
			projectDirPath,
		)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// config.
		err = services.Project.CreateDefautlFile(
			config.ConfigFile,
			"config.yaml",
			configFileName,
			projectDirPath,
		)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = services.Project.CreateDefautlFile(
			config.ConfigFile,
			"load_balancer.conf",
			"load_balancer.conf",
			fileDirPath,
		)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create .nomad.hcl configuration file.
		// Parse chart config file.
		chartFilePath := filepath.Join(projectDirPath, chartFileName)
		chartFile, err := os.ReadFile(chartFilePath)
		if err != nil {
			fmt.Printf("failed to read chart file, %s\n", err)
			os.Exit(1)
		}

		parsedChartConfig, err := services.Parser.ParseYAML(chartFile)
		if err != nil {
			fmt.Printf("failed to parsing chart file, %s\n", err)
			os.Exit(1)
		}

		chartConfig := services.Parser.ParseConfig("chart", parsedChartConfig)

		// Parse job config file.
		jobFilePath := filepath.Join(projectDirPath, configFileName)
		jobFile, err := os.ReadFile(jobFilePath)
		if err != nil {
			fmt.Printf("failed to read job file, %s\n", err)
			os.Exit(1)
		}

		parsedJobConfig, err := services.Parser.ParseYAML(jobFile)
		if err != nil {
			fmt.Printf("failed to parsing job file, %s\n", err)
			os.Exit(1)
		}

		jobConfig := services.Parser.ParseConfig(
			"job",
			parsedJobConfig["job"].(map[string]interface{}),
		)

		configStructure := services.Builder.BuildConfigStructure(
			jobConfig,
			chartConfig,
			projectDirPath,
		)

		err = services.Output.CreateConfigFile(
			configName,
			projectDirPath,
			configStructure,
		)

		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Flags.
	initCmd.Flags().StringP("name", "n", "", "project name")
}
