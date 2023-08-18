package cmd

import (
	"fmt"
	"log"
	"os"
	"prism/config"

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
		projectDir := "prism-default"
		if name != "" {
			projectDir = name
		}

		projectFileDir := "files"
		projectPath := fmt.Sprintf("%s/%s", rootDir, projectDir)

		dirStat, err := os.Stat(fmt.Sprintf("%s/%s", rootDir, projectDir))
		if err != nil || !dirStat.IsDir() {
			err = os.MkdirAll(
				fmt.Sprintf("%s/%s", projectDir, projectFileDir),
				0700,
			)
			if err != nil {
				log.Fatalf("project initialization error, %s", err)
			}
		}

		// Create default project files.
		defaultFiles := []string{"prism.yaml", "default_config.yaml"}
		for _, file := range defaultFiles {
			err = services.Project.CreateDefautlFile(
				config.ConfigFile,
				file,
				projectPath,
			)

			if err != nil {
				log.Fatalln(err)
			}
		}

		fileDirPath := fmt.Sprintf("%s/%s", projectPath, projectFileDir)
		err = services.Project.CreateDefautlFile(
			config.ConfigFile,
			"load_balancer.conf",
			fileDirPath,
		)

		if err != nil {
			log.Fatalln(err)
		}

		// Create .nomad.hcl configuration file.
		// Parse chart config file.
		chartPath := fmt.Sprintf("%s/prism.yaml", projectPath)
		chartFile, err := os.ReadFile(chartPath)
		if err != nil {
			log.Fatalf("error read file, %s", err)
		}

		chartConfig, err := services.Parser.ParseChart(chartFile)
		if err != nil {
			log.Fatalln(err)
		}

		// Parse job config file.
		defaultConfigFile := fmt.Sprintf("%s/default_config.yaml", projectPath)
		file, err := os.ReadFile(defaultConfigFile)
		if err != nil {
			log.Fatalf("error read file, %s", err)
		}

		defaultConfig, err := services.Parser.ParseJobConfig(file)
		if err != nil {
			log.Fatalln(err)
		}

		template := services.Builder.BuildConfigTemplate(defaultConfig, chartConfig)

		err = services.Output.CreateNomadConfigFile(
			"default_config",
			projectPath,
			template,
		)
		if err != nil {
			log.Fatalln(err)
		}

		// printJob(configYAML)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Flags.
	initCmd.Flags().StringP("name", "n", "", "project name")
	initCmd.Flags().StringP("chart", "c", "", "path to chart file")
}

// func printJob(job model.ConfigBlock) {
// 	// fmt.Println(job.Name)
// 	fmt.Print("parameters:\n\n")
// 	for _, p := range job.Parameter {
// 		fmt.Printf("%v\n", p)
// 	}

// 	fmt.Print("\n\nblocks:\n\n")
// 	for _, b := range job.Block {
// 		fmt.Printf("%v\n\n", b)

// 		// 	if b.Name == "group" {
// 		// 		for _, b := range b.Block {
// 		// 			for _, v := range b.Parameter {
// 		// 				for k, v := range v {
// 		// 					fmt.Printf("%v: %v [%T]\n", k, v, v)
// 		// 				}
// 		// 			}
// 		// 		}
// 		// 	}
// 	}

// 	// fmt.Printf("\njob struct parameter's: %v\n", job.Parameter)
// 	// fmt.Printf("\njob struct block's: %v\n", job.Block)

// 	services.Builder.BuildConfigTemplate(job)
// }
