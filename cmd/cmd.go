package cmd

import (
	"fmt"
	"github.com/bidaya0/gbatect/converter"
	batecttypes "github.com/bidaya0/gbatect/types"
	"github.com/compose-spec/compose-go/loader"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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

		if fromfile != "" {
			dockercomposefile, err := os.ReadFile(fromfile)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
			k1, err := loader.ParseYAML(dockercomposefile)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
			tmpk := k1["services"]
			tmp3, _ := tmpk.(map[string]interface{})
			services, err := converter.LoadServices(tmp3)
			containers, err := converter.TransServicesToContainer(services)
			var f1 = batecttypes.BatectConfig{
				Containers: containers,
			}
			batectyaml, err := yaml.Marshal(&f1)
			if err != nil {
				fmt.Printf("error: %v", err)
			}

			if err != nil {
				fmt.Printf("error: %v", err)
			}
			if tofile != "" {
				err := os.WriteFile(tofile, batectyaml, 0644)
				if err != nil {
					fmt.Printf("%v", err)
				}
			} else {
				fmt.Printf("%v", string(batectyaml))
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
