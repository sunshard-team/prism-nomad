// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package deployment

import (
	"fmt"
	"prism/internal/model"
	"prism/internal/service/builder"
	"prism/internal/service/parser"
	"slices"

	"github.com/hashicorp/nomad/api"
)

type Deployment struct {
	parser  parser.Parser
	builder builder.StructureBuilder
	changes builder.Changes
}

func NewDeployment(
	parser parser.Parser,
	builder builder.StructureBuilder,
	changes builder.Changes,
) *Deployment {
	return &Deployment{
		parser:  parser,
		builder: builder,
		changes: changes,
	}
}

// Checks whether the namespace exists in the cluster.
// If the --create-namespace flag is specified and
// the specified namespace does not exist, then it will be created.
func (s *Deployment) CheckNamespace(namespace model.CheckNamespace) error {
	var namespaces []string
	client := namespace.Client

	availableNamespaces, _, err := client.Namespaces().List(&api.QueryOptions{})
	if err != nil {
		return err
	}

	for _, n := range availableNamespaces {
		namespaces = append(namespaces, n.Name)
	}

	if !slices.Contains(namespaces, namespace.Namespace) {
		if namespace.CreateNamespace {
			newNamespace := &api.Namespace{
				Name: namespace.Namespace,
			}

			_, err := client.Namespaces().Register(
				newNamespace,
				&api.WriteOptions{},
			)

			if err != nil {
				return fmt.Errorf("error create namespace %s", err)
			}

			fmt.Printf(
				"Namespace \"%s\" was successfully created.\n",
				namespace.Namespace,
			)

			return nil
		}

		return fmt.Errorf(
			"%s %s",
			"The specified namespace does not exist.",
			"Use the --create-namespace flag to create a new namespace.",
		)
	}

	return nil
}

// Job configuration deployment in the nomad cluster.
func (s *Deployment) Deployment(client *api.Client, config string) (string, error) {
	job, _ := client.Jobs().ParseHCL(config, true)

	_, _, err := client.Jobs().Register(job, &api.WriteOptions{})
	if err != nil {
		return "", fmt.Errorf("job registration error, %s", err)
	}

	return *job.ID, nil
}
