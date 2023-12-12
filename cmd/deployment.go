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

		CACert, err := cmd.Flags().GetString("ca-cert")
		if err != nil {
			fmt.Printf("failed to read flag \"ca-cert\", %s\n", err)
			os.Exit(1)
		}

		CAPath, err := cmd.Flags().GetString("ca-path")
		if err != nil {
			fmt.Printf("failed to read flag \"ca-path\", %s\n", err)
			os.Exit(1)
		}

		clientCert, err := cmd.Flags().GetString("client-cert")
		if err != nil {
			fmt.Printf("failed to read flag \"client-cert\", %s\n", err)
			os.Exit(1)
		}

		clientKey, err := cmd.Flags().GetString("client-key")
		if err != nil {
			fmt.Printf("failed to read flag \"client-key\", %s\n", err)
			os.Exit(1)
		}

		TLSServerName, err := cmd.Flags().GetString("tls-server-name")
		if err != nil {
			fmt.Printf("failed to read flag \"tls-server-name\", %s\n", err)
			os.Exit(1)
		}

		TLSSkipVerify, err := cmd.Flags().GetBool("tls-skip-verify")
		if err != nil {
			fmt.Printf("failed to read flag \"tls-skip-verify\", %s\n", err)
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

		var CACertPEM []byte
		var clientCertPEM []byte
		var clientKeyPEM []byte

		if CACert != "" {
			CACertPEM, err = os.ReadFile(CACert)
			if err != nil {
				fmt.Printf("error read cert, %s", err)
				os.Exit(1)
			}
		}

		if clientCert != "" {
			clientCertPEM, err = os.ReadFile(clientCert)
			if err != nil {
				fmt.Printf("error read cert, %s", err)
				os.Exit(1)
			}
		}

		if clientKey != "" {
			clientKeyPEM, err = os.ReadFile(clientKey)
			if err != nil {
				fmt.Printf("error read cert, %s", err)
				os.Exit(1)
			}
		}

		tlsConfigAPI := &api.TLSConfig{
			CACert:        CACert,
			CAPath:        CAPath,
			CACertPEM:     CACertPEM,
			ClientCert:    clientCert,
			ClientCertPEM: clientCertPEM,
			ClientKey:     clientKey,
			ClientKeyPEM:  clientKeyPEM,
			TLSServerName: TLSServerName,
			Insecure:      !TLSSkipVerify,
		}

		configAPI := &api.Config{
			Address:   address,
			SecretID:  token,
			TLSConfig: tlsConfigAPI,
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
			fmt.Printf("an error occurred while checking the namespace: %s\n", err)
			os.Exit(1)
		}

		jobID, err := services.Deployment.Deployment(client, outputConfig)
		if err != nil {
			fmt.Printf("an error occurred while deploying the job: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Job \"%s\" deployed successfully.\n", jobID)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("path", "p", "", "path to project directory") // required
	deployCmd.Flags().StringP("release", "r", "", "release name")
	deployCmd.Flags().StringP("namespace", "n", "", "namespace name")

	deployCmd.Flags().Bool(
		"create-namespace",
		false,
		"create a namespace in the cluster if one is not created",
	)

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

	deployCmd.Flags().StringP("address", "a", "", "cluster address") // required for deployment
	deployCmd.Flags().StringP("token", "t", "", "cluster access token")

	deployCmd.Flags().String(
		"ca-cert",
		"",
		"Path to a PEM encoded CA cert file to use to verify the Nomad server SSL certificate.",
	)

	deployCmd.Flags().String(
		"ca-path",
		"",
		"Path to a directory of PEM encoded CA cert files to verify the Nomad server SSL certificate.",
	)

	deployCmd.Flags().String(
		"client-cert",
		"",
		"Path to a PEM encoded client certificate for TLS authentication to the Nomad server.",
	)

	deployCmd.Flags().String(
		"client-key",
		"",
		"Path to an unencrypted PEM encoded private key matching the client certificate from --client-cert.",
	)

	deployCmd.Flags().String(
		"tls-server-name",
		"",
		"The server name to use as the SNI host when connecting via TLS.",
	)

	deployCmd.Flags().Bool(
		"tls-skip-verify",
		true,
		"Do not verify TLS certificate. This is highly not recommended.",
	)
}
