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
			log.Fatalf(`error read flag "name", %s`, err)
		}

		// Get root dir path.
		rootDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("project initialization error, %s", err)
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
				log.Fatalf("project initialization error, %s", err)
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
			log.Fatalln(err)
		}

		// config.
		err = services.Project.CreateDefautlFile(
			config.ConfigFile,
			"config.yaml",
			configFileName,
			projectDirPath,
		)

		if err != nil {
			log.Fatalln(err)
		}

		err = services.Project.CreateDefautlFile(
			config.ConfigFile,
			"load_balancer.conf",
			"load_balancer.conf",
			fileDirPath,
		)

		if err != nil {
			log.Fatalln(err)
		}

		// Create .nomad.hcl configuration file.
		// Parse chart config file.
		chartFilePath := filepath.Join(
			projectDirPath,
			chartFileName,
		)

		chartFile, err := os.ReadFile(chartFilePath)
		if err != nil {
			log.Fatalf("error read file, %s", err)
		}

		parsedChartConfig, err := services.Parser.ParseYAML(chartFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		chartConfig := services.Parser.ParseConfig("chart", parsedChartConfig)

		// Parse job config file.
		jobFilePath := filepath.Join(
			projectDirPath,
			configFileName,
		)

		jobFile, err := os.ReadFile(jobFilePath)
		if err != nil {
			log.Fatalf("error read file, %s", err)
		}

		parsedJobConfig, err := services.Parser.ParseYAML(jobFile)
		if err != nil {
			fmt.Println(err)
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
