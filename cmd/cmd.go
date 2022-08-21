package cmd

import (
	"fmt"
	"github.com/bidaya0/gbatect/converter"
	"github.com/spf13/cobra"
	"os"
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
	Short: "A batect file convertor",
	Long: `gbatect is a tool help users move exists docker-compose to batect.
	gbatect take the docker-compose.yml and translates it to batect.yml.
`,
	Run: func(cmd *cobra.Command, args []string) {
		//if len(args) > 0{
		//	filepath := args[0]
		//}
		if fromfile != "" {
			result := converter.ReadAndConvert(fromfile)
			if tofile != "" {
				err := os.WriteFile(tofile, result, 0644)
				if err != nil {
					fmt.Printf("%v", err)
				}
			} else {
				fmt.Printf("%v", string(result))
			}
		}
	},
}

// Execute executes the root command.
func Execute() error {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&fromfile, "fromfile", "f", "", "Source directory to read from")
	convertCmd.Flags().StringVarP(&tofile, "tofile", "t", "", "Target directory to output")
	return rootCmd.Execute()
}

func init() {
}
