# Prism

[![Discord](https://img.shields.io/badge/prism-cli.svg?style=flat&logo=discord)](https://discord.gg/fSvtfPTrud)
[![Telegram](https://img.shields.io/badge/Telegram-Join%20Chat-blue?logo=telegram)](https://t.me/+Ubx2ygV2rd4yNzUy)

Prism is a tool that simplifies the creation of Nomad job configuration templates and deploys them to a remote cluster.

![Scheme of work Prism cli](docs/prism.svg)

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Flags](#flags)
- [Example command](#example-command)
- [Pack Information](#pack-information)

## Prerequisites

+ Go >= 1.21.1 is [installed](https://go.dev/doc/install)
+ Nomad is [installed](https://developer.hashicorp.com/nomad/tutorials/get-started/gs-install)

## Installation

**To get started with Prism, you'll need to install it. Please follow these steps:**

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/sunshard-prism/prism-nomad
   ```

2. Create and move the prism binary to /usr/local/bin:
   ```bash
   cd prism-nomad && make build-prism
   ```

3. Move the prism binary to /usr/local/bin:
   ```bash
   mv build/prism /usr/local/bin/prism
   ```

4. Grants executable permissions for prism binary:
   ```bash
   chmod +x /usr/local/bin/prism
   ```

5. Test to ensure the version you installed is up-to-date:
   ```bash
   prism --version
   ```

**Or download pre-built binary (Windows, MacOS or Linux).**

[Release download](https://github.com/sunshard-prism/prism-nomad/releases)

## Usage

**1. Creating a project. In prism they are called “pack”.**

   To do this, run the command, where \<name> is the name of your project (default name pack name "prism", further in the example this name will be indicated).
   ```
   prism init <name>
   ```

   Project directory and default files will be created:
   - pack.yaml (required) - details [Pack Information](#pack-information)
   - config.yaml (required) - nomad job configuration
   - files directory - directory for additional files

**2. Creating Nomad job configuration.**

   Now you need to describe the job configuration in the config.yml file. Describe the basic settings and configuration blocks (that can be reused, such as common settings for production and development environments) of your project.

**3. Creating additional files.**

   Next, create additional files in the files directory (unnecessary files and files created by default can be deleted).

   Let's look at the example of production and work environments. Create two files dev.yaml and prod.yaml.
   For each environment, describe the parameters and configuration blocks.

**4. Deploying Nomad configuration job.**

   And so, the main file (config.yaml) and additional job configuration (from example dev.yaml and prod.yaml) files are described, now we are ready to make the first deployment of the job.

   Let's deploy the task in the development environment.
   First, we’ll do a dry run to output the finished configuration in HCL format to make sure that everything is specified correctly.
   You can also output the result to a file in HCL format using the `--output` flag.

   ```
   prism deploy --path ./prism --release dev --file dev.yaml --dry-run
   ```

   *Note that we add `--file dev.yaml` to tell prism to take the dev.yaml file and overlay it with the main config file. Prism adds parameters and configuration blocks missing from the main configuration file or replaces them if they exist (taking into account the hierarchy and the official specification of nomad). If multiple files are specified, changes will be applied in the order in which the files are specified.*

   If everything is in order, we can deploy the job configuration.

   ```
   prism deploy --path ./prism --release dev --file dev.yaml --address $cluster-address
   ```

   *If a token is required to access the cluster, simply specify it by adding the `--token` flag.*

   Ready! The prod version (using prod.yaml) can be deployed in the same way as the dev version.

## Commands

Prism provides the following commands:

- `init`: Create a new project.
- `deploy`: Deploy a configuration to a remote cluster.

For more details on each command and their usage, run `prism [command] --help`.

## Flags

**prism:**
- `--version`: Use this flag to request the current version.

**init command:**
- For init command use project name argument `prism init <name>`.

**deploy command:**
- `-a, --address string`: The address of the Nomad cluster.
- `-t, --token string`: Cluster access token.
- `-n, --namespace string`: Namespace name.
- `-r, --release string`: Release name.
- `-p, --path string`: Path to the project directory.
- `-o, --output string`: Path to the directory where the `<project>_<release>.nomad.hcl` file will be created.
- `-f, --file strings`: File name or full path to the file to update the configuration.
- `--create-namespace`: Create a namespace in the cluster if it doesn't exist.
- `--dry-run`: Print the job configuration to the console (blocking the deployment).

## Example command:

Here's an example of deploying a configuration to a remote Nomad cluster:

```bash
prism deploy -a http://nomad_ip:4646 -t nomad_token -n dest_namespace -r name_of_release -p /path/to/prismpack 
```

This command will perform a dry run and print the job configuration to the console. Adjust the flags to suit your deployment needs.

## Pack Information

The `pack.yaml` file is used in the context of Prism-cli packages, which serve as a way to describe, package, and deploy applications in Nomad.
This file contains metadata and information about the Prism Pack, which is an archive containing descriptions of Nomad resources, default values for creating deployed applications in the Nomad cluster.

Here are some of the key fields that may be found in the `pack.yaml` file:

- **name**: The name of the Prism Pack.
- **description**: Description of the Prism Pack.
- **maintainers**: Information about those who maintain the Prism Pack.
- **type**: Specifies the Nomad scheduler to use. Nomad provides the service, system, batch, and sysbatch schedulers.
- **sources**: Links to source code or resources associated with the Prism Pack.
- **deploy_version**: The version of the application it contains.
- **prism_version**: The version of the Prism Pack, aiding in tracking changes and updates to the Prism Pack.
- **nomad_version**: The version of Nomad on which the Prism Pack has been tested.
- **dependencies**: Specifies dependencies of the current Prism Pack on other Prism Packs, which will be automatically installed when installing the main Prism Pack.

This file is valuable for organizing and documenting Prism Packs, as well as for their publication and exchange among Nomad developers.
The `pack.yaml` file helps manage Prism Pack versions, simplifies searching and describing packs, and eases their utilization in the Nomad environment.