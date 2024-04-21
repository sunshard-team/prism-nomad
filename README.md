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
- [Pack information](#pack-information)
- [Environment variables](#environment-variables)
- [Pack dependencies](#pack-dependencies)
- [Deployment status](#deployment-status)
- [Release](#release)
- [Sidecar service](#sidecar-service)

## Prerequisites

+ Go >= 1.21.5 is [installed](https://go.dev/doc/install)
+ Nomad is [installed](https://developer.hashicorp.com/nomad/tutorials/get-started/gs-install)

## Installation

**To get started with Prism, you'll need to install it. Please follow these steps:**

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/sunshard-team/prism-nomad
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

[Release download](https://github.com/sunshard-team/prism-nomad/releases)

## Usage

**1. Creating a project. In prism they are called “pack”.**

   To do this, run the command, where \<name> is the name of your project (default name pack name "prism", further in the example this name will be indicated).

   ```bash
   prism init <name>
   ```

   Project directory and default files will be created:
   - pack.yaml (required) - details [Pack Information](#pack-information)
   - config.yaml (required) - nomad job configuration
   - files directory - directory for additional files

**2. Creating Nomad job configuration.**

   Now you need to describe the job configuration in the config.yaml file. Describe the basic settings and configuration blocks (that can be reused, such as common settings for production and development environments) of your project.

**3. Creating additional files.**

   Next, create additional files in the files directory (unnecessary files and files created by default can be deleted).

   Let's look at the example of production and work environments. Create two files dev.yaml and prod.yaml.
   For each environment, describe the parameters and configuration blocks.

**4. Deploying Nomad configuration job.**

   And so, the main file (config.yaml) and additional job configuration (from example dev.yaml and prod.yaml) files are described, now we are ready to make the first deployment of the job.

   Let's deploy the task in the development environment.
   First, we’ll do a dry run to output the finished configuration in HCL format to make sure that everything is specified correctly.
   You can also output the result to a file in HCL format using the `--output` flag.

   ```bash
   prism deploy --path ./prism --release dev --file dev.yaml --dry-run
   ```

   *Note that we add `--file dev.yaml` to tell prism to take the dev.yaml file and overlay it with the main config file. Prism adds parameters and configuration blocks missing from the main configuration file or replaces them if they exist (taking into account the hierarchy and the official specification of nomad). If multiple files are specified, changes will be applied in the order in which the files are specified.*

   If everything is in order, we can deploy the job configuration.

   ```bash
   prism deploy --path ./prism --release dev --file dev.yaml --address $cluster-address
   ```

   *If a token is required to access the cluster, simply specify it by adding the `--token` flag.*

   Ready! The prod version (using prod.yaml) can be deployed in the same way as the dev version.

## Commands

   Prism provides the following commands:

   - `init`: Create a new project.
   - `deploy`: Deploy a configuration to a remote cluster.
      - `tls`: Parameters required to configure TLS on the HTTP client used to communicate with Nomad.

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
   - `-w, --wait-time`: Deployment wait time in seconds (default 120 sec.).
   - `-e, --env`: Environment variables in the form key=value.
   - `--env-file`: Full path to the file with environment variables.
   - `--create-namespace`: Create a namespace in the cluster if it doesn't exist.
   - `--dry-run`: Print the job configuration to the console (blocking the deployment).
   
   **tls command:**
   - `--ca-cert`: Path to a PEM encoded CA cert file to use to verify the Nomad server SSL certificate.
   - `--ca-path`: Path to a directory of PEM encoded CA cert files to verify the Nomad server SSL certificate.
   - `--client-cert`: Path to a PEM encoded client certificate for TLS authentication to the Nomad server.
   - `--client-key`: Path to an unencrypted PEM encoded private key matching the client certificate from --client-cert.
   - `--tls-server-name`: The server name to use as the SNI host when connecting via TLS.
   - `--tls-skip-verify`: Do not verify TLS certificate. This is highly not recommended.

## Example command:

   Here's an example of deploying a configuration to a remote Nomad cluster:

   ```bash
   prism deploy -a http://nomad_ip:4646 -t nomad_token -n dest_namespace -r name_of_release -p /path/to/prismpack 
   ```

   This command will perform a dry run and print the job configuration to the console. Adjust the flags to suit your deployment needs.

## Pack information

   The `pack.yaml` file is used in the context of Prism-cli packages, which serve as a way to describe, package, and deploy applications in Nomad. This file contains metadata and information about the Prism Pack, which is an archive containing descriptions of Nomad resources, default values for creating deployed applications in the Nomad cluster.

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

   This file is valuable for organizing and documenting Prism Packs, as well as for their publication and exchange among Nomad developers. The `pack.yaml` file helps manage Prism Pack versions, simplifies searching and describing packs, and eases their utilization in the Nomad environment.

## Environment variables

   In all the files from which the job Nomad will ultimately be built (configuration file, additional template files, etc.), local environment variables can be added in the parameters and block labels.

   ### What it looks like and how it works.

   **Syntax:** `"${PRISM_ANY_VARIABLE_NAME}"`\
   Example: `"${PRISM_GROUP_COUNT}"`

   **Rule:**\
   The variable name must begin with the prefix `PRISM_` (as indicated - with an underscore). The variable must be specified as a string, i.e. in double quotes "" (if it is already in the string, there is no need to wrap it in double quotes), inside curly braces `{}`, preceded by a dollar sign `$`, i.e. overall it looks like `"${PRISM_VAR}"`.

   **How it works.**\
   For example, let's specify a variable for the job name and a count for the group:

   ```yaml
   job:
     name: "${PRISM_JOB_NAME}"
     ...

     group:
       - name: "group-name"
         count: "${PRISM_GROUP_COUNT}"
         ...
   ```

   The value of a variable can be taken from three sources:
   1. **low priority** - Local environment variables
   2. **medium priority** - File with variables (yaml, toml, json, etc.)
   3. **high priority** - Flag `--env` which accepts `key=value`

   If the same variable is specified in 2 or all 3 sources, its value will be taken from the most priority source, i.e. When deploying, you can additionally specify a file with variables and add a value through a special flag.

   - If you need to take the value of variables only from the local environment, then during deployment there is no need to specify any additional parameters (flags). Simply specify the necessary variables in the configuration file and additional template files.

   - If you need to take the value of a variable specified in a file, when deploying, specify the full path to the file with the variables (including file name), using the `--env-file` flag.

   - If you want to specify the value of a variable in the deployment command, specify the key (variable name) and value as `key=value` using the `--env` flag. Example: `--env PRISM_JOB_NAME=job-name`. To specify multiple variables, simply write them separated by commas (without a space after the comma), example: `--env PRISM_JOB_NAME=job-name,PRISM_GROUP_COUNT=2`.

   The number of variables in one line is not limited, i.e. You can specify, for example, the following line `"../path/${PRISM_ANY_VAR}/path/${PRISM_ANY_VAR_TWO}/"`.

   ### Default value.

   You can specify a default value for an environment variable. It will be taken if the variable is not found in any of the sources. It is indicated immediately after the variable name, separated by a vertical bar with the keyword "default=", example: `${PRISM_VAR|default=any-value}`.

   **Do not leave the default without a value `"${PRISM_VAR|default=}"`, otherwise the line will be ignored!**

## Pack dependencies

   You can specify dependencies for a Pack to deploy them sequentially, before deploying the main job. A dependency is any other package, or rather its “basic” job configuration template - `config.yaml` file.

   Dependency parameters:
   - `name`: Dependency name.
   - `pack_version`: Pack vesrion (optional).
   - `path`: Full path to the Pack directory.
   - `files`: List of files name or full paths to files to update (parameter overrides/additions), configuration. If only the filename is specified, Prism will look for it in the current Pack rather than the dependency Pack. This works like the `--file` flag of the `deploy` command.

   The jobs is deployed in the following order:
   1. Dependencies deployment, in the order in which they are listed;
   2. Jobs deployment from the current Pack (job for which the dependencies are indicated);

   The `--dry-run` flag prints jobs to the console in the order in which they will be deployed.

   When jobs are deployed, the deployment status will be displayed in the console, [deployment status](#deployment-status).

## Deployment status

   Starting with version v0.4.0, the job deployment status functionality is introduced.

   When deploying jobs, the following statuses are displayed in the console: deployment, job, allocation and deployment time of each job. If an error occurs during the deployment process, the process will be stopped.

   Additionally, a wait time of 2 minutes is set for the deployment of each job. You can change the waiting time for jobs to be deployed using the `--wait-time` flag (the time is indicated in seconds).
   
   **The job will be considered successfully deployed only if the deployment status is "successful"!**

## Release

   During deployment, you can specify any release name. It allows you to deploy one job under different releases, using the `--release` flag. Starting from version v0.4.0, when specifying a release, it will be added by default to the name of `job`, `group`, `task`, `device`.


## Sidecar service

   To specify the default `sidecar_service` value:

   ```
   connect {
     sidecar_service {}
   }
   ```

   use the `open_sidecar_service` parameter with the value `true`:

   ```yaml
   connect:
     open_sidecar_service: true
   ```

   Otherwise, the parameter need not be specified. If you decide to leave it in place, simply set it to `false`, (in any case, this parameter will be removed from the final configuration).
