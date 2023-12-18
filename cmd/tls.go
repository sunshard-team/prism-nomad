// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"os"

	"github.com/hashicorp/nomad/api"
	"github.com/spf13/cobra"
)

var TLSCmd = &cobra.Command{
	Use:   "tls",
	Short: "Parameters required to configure TLS.",
	Long:  "Parameters required to configure TLS on the HTTP client used to communicate with Nomad.",
	Run: func(cmd *cobra.Command, args []string) {
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

		TLSConfigAPI = &api.TLSConfig{
			CACert:        CACert,
			CAPath:        CAPath,
			CACertPEM:     CACertPEM,
			ClientCert:    clientCert,
			ClientCertPEM: clientCertPEM,
			ClientKey:     clientKey,
			ClientKeyPEM:  clientKeyPEM,
			TLSServerName: TLSServerName,
			Insecure:      TLSSkipVerify,
		}

		deployment(cmd, args)
	},
}

func init() {
	deployCmd.AddCommand(TLSCmd)

	TLSCmd.Flags().String(
		"ca-cert",
		"",
		"Path to a PEM encoded CA cert file to use to verify the Nomad server SSL certificate.",
	)

	TLSCmd.Flags().String(
		"ca-path",
		"",
		"Path to a directory of PEM encoded CA cert files to verify the Nomad server SSL certificate.",
	)

	TLSCmd.Flags().String(
		"client-cert",
		"",
		"Path to a PEM encoded client certificate for TLS authentication to the Nomad server.",
	)

	TLSCmd.Flags().String(
		"client-key",
		"",
		"Path to an unencrypted PEM encoded private key matching the client certificate from --client-cert.",
	)

	TLSCmd.Flags().String(
		"tls-server-name",
		"",
		"The server name to use as the SNI host when connecting via TLS.",
	)

	TLSCmd.Flags().Bool(
		"tls-skip-verify",
		true,
		"Do not verify TLS certificate. This is highly not recommended.",
	)
}
