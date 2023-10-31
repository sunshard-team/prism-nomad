# Prism

[![Discord](https://img.shields.io/badge/prism-cli.svg?style=flat&logo=discord)](https://discord.gg/fSvtfPTrud)

Prism is a tool that simplifies the creation of Nomad job configuration templates and deploys them to a remote cluster.

![Scheme of work Prism cli](docs/prism.svg)

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Flags](#flags)
- [Example command](#example-command)
- [Pack Information](#pack-information)

## Installation

**To get started with Prism, you'll need to install it. Please follow these steps:**

1. Download release on local station:
   ```bash
   wget https://github.com/sunshard-prism/prism-nomad/releases/download/{{version}}/prism.linux-amd64.tar.gz

2. Move prism binaries to /usr/local/bin:
   ```bash
   mv prism /usr/local/bin/prism
   ```

3. Grants executable permissions for prism binaries:
   ```bash
   chmod +x /usr/local/bin/prism
   ```

4. Test to ensure the version you installed is up-to-date:
   ```bash
   prism —version
   ```
   
**Or download pre-built binary (Windows, MacOS, or Linux).**

[Release download](https://github.com/sunshard-prism/prism-nomad/releases)

## Usage

Prism simplifies the process of creating and deploying Nomad job configurations. You can define your infrastructure and application requirements in a `config.yaml` file and then generate configuration files and Go code for deployment.

### Flags

To deploy a configuration to a remote Nomad cluster, use the `deploy` command. Here's how to use it:

```bash
prism deploy [flags]
```

## Commands

Prism provides the following commands:

- `deploy`: Deploy a configuration to a remote cluster.
- `init`: Create a new project.

For more details on each command and their usage, run `prism [command] --help`.

**Flags**:

- `-a, --address string`: The address of the Nomad cluster.
- `--create-namespace`: Create a namespace in the cluster if it doesn't exist.
- `--dry-run`: Print the job configuration to the console (blocking the deployment).
- `-f, --file strings`: File name or full path to the file to update the configuration.
- `-n, --namespace string`: Namespace name.
- `-o, --output string`: Path to the directory where the `<project>_<release>.nomad.hcl` file will be created.
- `-p, --path string`: Path to the project directory.
- `-r, --release string`: Release name.
- `-t, --token string`: Cluster access token.
- `--version`: Use this flag to request the current version.

### Example command:

Here's an example of deploying a configuration to a remote Nomad cluster:

```bash
prism deploy -a http://nomad_ip:4646 -t nomad_token -n dest_namespace -r name_of_release -p /path/to/prismpack 
```

This command will perform a dry run and print the job configuration to the console. Adjust the flags to suit your deployment needs.

## Pack Information

The `pack.yaml` file is used in the context of Prism-cli packages, which serve as a way to describe, package, and deploy applications in Nomad. This file contains metadata and information about the Prism Pack, which is an archive containing descriptions of Nomad resources, default values for creating deployed applications in the Nomad cluster. Here are some of the key fields that may be found in the `pack.yaml` file:

- **name**: The name of the Prism Pack.
- **description**: Description of the Prism Pack.
- **maintainers**: Information about those who maintain the Prism Pack.
- **type**: Specifies the Nomad scheduler to use. Nomad provides the service, system, batch, and sysbatch schedulers.
- **sources**: Links to source code or resources associated with the Prism Pack.
- **deploy_version**: The version of the application it contains.
- **prism_version**: The version of the Prism Pack, aiding in tracking changes and updates to the Prism Pack.
- **nomad_version**: The version of Nomad on which the Prism Pack has been tested.
- **dependencies**: Specifies dependencies of the current Prism Pack on other Prism Packs, which will be automatically installed when installing the main Prism Pack.

This file is valuable for organizing and documenting Prism Packs, as well as for their publication and exchange among Nomad developers. The `pack.yaml` file helps manage Prism Pack versions, simplifies searching and describing packs, and eases their utilization in the Nomad environment.