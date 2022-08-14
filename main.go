package main

import (
	"fmt"
	"github.com/compose-spec/compose-go/loader"
//	"github.com/compose-spec/compose-go/types"
	"gopkg.in/yaml.v3"
	"os"
//	"strings"
	converter "github.com/bidaya0/gbatect/converter"
	batecttypes "github.com/bidaya0/gbatect/types"	
)


func main() {
	dockercomposefile, err := os.ReadFile("./docker-compose.yml")
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
	var f1 batecttypes.BatectConfig 
	f1.Containers = containers
	batectyaml, err := yaml.Marshal(&f1)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	os.WriteFile("./batect.yml", batectyaml, 0666)
}
