package cmd

import (
  "github.com/bidaya0/gbatect/converter"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "gbatect",
		Short: "A batect file convertor",
		Long: `gbatect is a tool help users move exists docker-compose to batect.
		gbatect take the docker-compose.yml and translates it to batect.yml.
`,
    Run: func(cmd *cobra.Command, args []string) {
				if len(args) > 0{
					filepath := args[0]
					converter.ReadAndConvert(filepath)
				}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
}

