package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var installLongDescription = fmt.Sprintf(
	"%s\n%s",
	"Deploying a configuration on a remote cluster,",
	"or outputting the configuration to the console or file.",
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Deploying a configuration to a remote cluster",
	Long:  installLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Printf("failed to read flag \"path\", %s\n", err)
			os.Exit(1)
		}

		path = filepath.Join(path)

		namespace, err := cmd.Flags().GetString("namespace")
		if err != nil {
			fmt.Printf("failed to read flag \"namespace\", %s\n", err)
			os.Exit(1)
		}

		release, err := cmd.Flags().GetString("release")
		if err != nil {
			fmt.Printf("failed to read flag \"release\", %s\n", err)
			os.Exit(1)
		}

		file, err := cmd.Flags().GetStringSlice("file")
		if err != nil {
			fmt.Printf("failed to read flag \"output\", %s\n", err)
			os.Exit(1)
		}

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			fmt.Printf("failed to read flag \"dry-run\", %s\n", err)
			os.Exit(1)
		}

		outputPath, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("failed to read flag \"output\", %s\n", err)
			os.Exit(1)
		}

		// address, err := cmd.Flags().GetString("address")
		// if err != nil {
		// 	fmt.Printf("failed to read flag \"address\", %s\n", err)
		// 	os.Exit(1)
		// }

		// token, err := cmd.Flags().GetString("token")
		// if err != nil {
		// 	fmt.Printf("failed to read flag \"token\", %s\n", err)
		// 	os.Exit(1)
		// }

		// createNamespace, err := cmd.Flags().GetBool("create-namespace")
		// if err != nil {
		// 	fmt.Printf("failed to read flag \"create-namespace\", %s\n", err)
		// 	os.Exit(1)
		// }

		if path == "" || namespace == "" {
			fmt.Printf(
				"%s %s %s\n",
				"failed execute install command,",
				"one of the required flags is not specified:",
				"path, version, namespace",
			)

			os.Exit(1)
		}

		// Get the project directory name.
		dirFormat, err := regexp.Compile(`.+\/(\w+\S+\w+)$`)
		if err != nil {
			fmt.Printf(
				"%s %s\n",
				"error execute install command,",
				"failed get project directory path",
			)

			os.Exit(1)
		}

		findProjectDir := dirFormat.FindStringSubmatch(path)
		projectDir := findProjectDir[1]

		// Create a configuration structure.
		parameter := model.ConfigParameter{
			ProjectDirPath: path,
			ProjectDir:     projectDir,
			Namespace:      namespace,
			Release:        release,
			Files:          file,
		}

		configStructure, err := services.Deployment.CreateConfigStructure(
			parameter,
		)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Dry run.
		if dryRun {
			if outputPath != "" {
				projectName := strings.ReplaceAll(projectDir, "-", "_")
				fileName := fmt.Sprintf("%s_config", projectName)

				err := services.Output.CreateConfigFile(
					fileName, outputPath, configStructure,
				)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			outputConfig, err := services.Output.OutputConfig(configStructure)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("\nOutput config:\n\n%v\n", outputConfig)
			return
		}

		// Deployment.
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringP("path", "p", "", "path to project directory") // required
	installCmd.Flags().StringP("namespace", "n", "", "namespace name")       // required
	installCmd.Flags().StringP("release", "r", "", "release name")
	installCmd.Flags().StringP("address", "a", "", "cluster address")
	installCmd.Flags().StringP("token", "t", "", "cluster access token")

	installCmd.Flags().StringSliceP(
		"file",
		"f",
		[]string{},
		"full path to file to update configuration",
	)

	installCmd.Flags().Bool(
		"create-namespace",
		false,
		"create namespace if not created",
	)

	installCmd.Flags().Bool(
		"dry-run",
		false,
		"output the result to the console (blocking the deployment)",
	)

	installCmd.Flags().StringP(
		"output",
		"o",
		"",
		fmt.Sprintf(
			`Path to the directory in which the "%s" file will be created`,
			"<project>_<release>_config_output.nomad.hcl",
		),
	)
}
