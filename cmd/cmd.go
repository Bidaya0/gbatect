package cmd

import (
	"github.com/bidaya0/gbatect/converter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gbatect",
	Short: "A batect file convertor",
	Long: `gbatect is a tool help users move exists docker-compose to batect.
	gbatect take the docker-compose.yml and translates it to batect.yml.
`,
}

var fromfile string
var tofile string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert docker-compose file to batect format",
	Long:  `take the docker-compose.yml and translates it to batect.yml.`,
	Run: func(cmd *cobra.Command, args []string) {
		converter.ConvertFiletoFile(fromfile, tofile)
	},
}

// Execute executes the root command.
func Execute() error {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&fromfile, "fromfile", "f", "", "Source directory to read from")
	convertCmd.Flags().StringVarP(&tofile, "tofile", "t", "/dev/stdout", "Target directory to output")
	return rootCmd.Execute()
}

func init() {
}
