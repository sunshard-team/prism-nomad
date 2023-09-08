package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "output",
	Short: "Create nomad configuration file",
	Long:  `Create nomad template configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf(`error read flag "name", %s`, err)
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf(`error read flag "path", %s`, err)
		}

		from, err := cmd.Flags().GetString("from")
		if err != nil {
			log.Fatalf(`error read flag "from", %s`, err)
		}

		chart, err := cmd.Flags().GetString("chart")
		if err != nil {
			log.Fatalf(`error read flag "chart", %s`, err)
		}

		if name != "" && path != "" {
			var output = services.Output
			var parser = services.Parser
			var builder = services.Builder

			// Parse chart config file.
			var chartConfig map[string]interface{}

			if chart != "" {
				chartfile, err := os.ReadFile(chart)
				if err != nil {
					log.Fatalf("error read file, %s", err)
				}

				chartConfig, err = parser.ParseChart(chartfile)
				if err != nil {
					log.Fatalf("error read file, %s", err)
				}
			}

			// Parse job config file.
			jobfile, err := os.ReadFile(from)
			if err != nil {
				log.Fatalf("error read file, %s", err)
			}

			jobConfig, err := parser.ParseJob(jobfile)
			if err != nil {
				log.Fatalln(err)
			}

			template := builder.BuildConfigStructure(jobConfig, chartConfig, "")

			_, err = output.OutputConfig(name, path, true, template)
			if err != nil {
				log.Fatalln(err)
			}

			return
		}

		fmt.Println("file name and path must be specified")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "name output file")
	createCmd.Flags().StringP(
		"path", "p", "",
		"directory where the configuration \"nomad\" file will be created",
	)
	createCmd.Flags().StringP(
		"from", "f", "", "path to configuration \"yaml\" file",
	)
	createCmd.Flags().StringP("chart", "c", "", "path to chart file")
}
