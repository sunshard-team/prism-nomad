// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"regexp"
	"strings"

	"github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

var deployLongDescription = fmt.Sprintf(
	"%s\n%s",
	"Deploying a configuration on a remote cluster,",
	"or outputting the configuration to the console or file.",
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploying a configuration to a remote cluster",
	Long:  deployLongDescription,
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

		address, err := cmd.Flags().GetString("address")
		if err != nil {
			fmt.Printf("failed to read flag \"address\", %s\n", err)
			os.Exit(1)
		}

		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Printf("failed to read flag \"token\", %s\n", err)
			os.Exit(1)
		}

		createNamespace, err := cmd.Flags().GetBool("create-namespace")
		if err != nil {
			fmt.Printf("failed to read flag \"create-namespace\", %s\n", err)
			os.Exit(1)
		}

		envFilePath, err := cmd.Flags().GetString("env-file")
		if err != nil {
			fmt.Printf("failed to read flag \"env-file\", %s\n", err)
			os.Exit(1)
		}

		envVars, err := cmd.Flags().GetStringToString("env")
		if err != nil {
			fmt.Printf("failed to read flag \"env\", %s\n", err)
			os.Exit(1)
		}

		if path == "" {
			fmt.Printf(
				"%s %s %s\n",
				"failed execute deploy command,",
				"one of the required flags is not specified:",
				"path",
			)

			os.Exit(1)
		}

		// Get the project directory name.
		dirFormat, err := regexp.Compile(`([\w+-]+)$`)
		if err != nil {
			fmt.Printf(
				"%s %s\n",
				"error execute deploy command,",
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
			EnvFilePath:    envFilePath,
			EnvVars:        envVars,
		}

		configStructure, err := services.Deployment.CreateConfigStructure(
			parameter,
		)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		outputConfig, err := services.Output.OutputConfig(configStructure)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Dry run.
		if dryRun {
			if outputPath != "" {
				projectName := strings.ReplaceAll(projectDir, "-", "_")
				fileName := projectName

				if release != "" {
					fileName = fmt.Sprintf("%s_%s", projectName, release)
				}

				err := services.Output.CreateConfigFile(
					fileName, outputPath, configStructure,
				)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			fmt.Printf("\nOutput config:\n\n%v\n", outputConfig)
			return
		}

		// Deployment.
		if address == "" {
			fmt.Printf(
				"%s %s %s\n",
				"failed execute deploy command,",
				"one of the required flags is not specified:",
				"address",
			)

			os.Exit(1)
		}

		configAPI := &api.Config{
			Address:  address,
			SecretID: token,
		}

		client, err := api.NewClient(configAPI)
		if err != nil {
			fmt.Printf("error create nomad api client %s", err)
			os.Exit(1)
		}

		checkNamespace := model.CheckNamespace{
			Client:          client,
			Namespace:       namespace,
			CreateNamespace: createNamespace,
		}

		err = services.Deployment.CheckNamespace(checkNamespace)
		if err != nil {
			fmt.Printf("namespace error: %s\n", err)
			os.Exit(1)
		}

		jobID, err := services.Deployment.Deployment(client, outputConfig)
		if err != nil {
			fmt.Printf("error job deployment: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Job \"%s\" deployed successfully.\n", jobID)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("address", "a", "", "cluster address")        // required for deployment
	deployCmd.Flags().StringP("token", "t", "", "cluster access token")     // required for deployment
	deployCmd.Flags().StringP("path", "p", "", "path to project directory") // required
	deployCmd.Flags().StringP("namespace", "n", "", "namespace name")
	deployCmd.Flags().StringP("release", "r", "", "release name")

	deployCmd.Flags().String(
		"env-file", "", "full path to the file with environment variables",
	)

	deployCmd.Flags().StringToStringP(
		"env", "e", map[string]string{}, "environment variables in the form key=value",
	)

	deployCmd.Flags().StringSliceP(
		"file",
		"f",
		[]string{},
		"file name or full path to file to update configuration",
	)

	deployCmd.Flags().Bool(
		"create-namespace",
		false,
		"create a namespace in the cluster if one is not created",
	)

	deployCmd.Flags().Bool(
		"dry-run",
		false,
		"print the job configuration to the console (blocking the deployment)",
	)

	deployCmd.Flags().StringP(
		"output",
		"o",
		"",
		fmt.Sprintf(
			"Path to the directory in which the \"%s\" file will be created",
			"<project>_<release>.nomad.hcl",
		),
	)
}
