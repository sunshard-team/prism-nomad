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
		chartPath := filepath.Join(
			projectDirPath,
			chartFileName,
		)

		chartFile, err := os.ReadFile(chartPath)
		if err != nil {
			log.Fatalf("error read file, %s", err)
		}

		chartConfig, err := services.Parser.ParseChart(chartFile)
		if err != nil {
			log.Fatalln(err)
		}

		// Parse job config file.
		configFile := filepath.Join(
			projectDirPath,
			configFileName,
		)

		file, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatalf("error read file, %s", err)
		}

		jobConfig, err := services.Parser.ParseJob(file)
		if err != nil {
			log.Fatalln(err)
		}

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
