package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
		"github.com/compose-spec/compose-go/types"
		"github.com/compose-spec/compose-go/loader"
		"strings"
		"os"
)

type ContainerOption struct {
	Image string ",omitempty"
	BuildDirectory string "build_directory,omitempty"
	Dockerfile string ",omitempty"
	Entrypoint string ",omitempty"
	Ports []string  ",omitempty"
	Volumes []string  ",omitempty"
	Environment types.MappingWithEquals  ",omitempty"
	BuildArgs map[string]string "build_args,omitempty" 
}

type ContainerItem struct {
    Name string
    Options map[string]ContainerOption
}

type Containers map[string]ContainerOption


func LoadServices(servicesDict map[string]interface{})([]types.ServiceConfig, error){
	var services []types.ServiceConfig
	for name ,element := range servicesDict {
		serviceConfig := &types.ServiceConfig{
			Name: name,
			Scale: 1,
		}
		loader.Transform(element, serviceConfig)
		services = append(services, *serviceConfig)
	}
	return services, nil
}

func TransServicesToContainer(servicesconfigs []types.ServiceConfig)(Containers, error){
	containers := make(map[string]ContainerOption)
	for _, service := range servicesconfigs {
			containeroption := ContainerOption{
			Image: service.Image,
			Environment: service.Environment,
			  }
			if service.Build != nil{
				containeroption.BuildDirectory = service.Build.Context
				containeroption.Dockerfile = service.Build.Dockerfile
			}
			if len(service.Entrypoint) > 0 {
				containeroption.Entrypoint = strings.Join(service.Entrypoint, " ")			
			}
			for _, port := range service.Ports{
				portstring := fmt.Sprintf("%v:%v", port.Published, port.Target)
				containeroption.Ports = append(containeroption.Ports,portstring)
			}
			for _, volume := range service.Volumes{
				volumestring := fmt.Sprintf("%v:%v", volume.Source, volume.Target)
				containeroption.Volumes = append(containeroption.Volumes,volumestring)
			}	

			containers[service.Name] = containeroption 
	}
	return containers, nil
}

func main() {
		dockercomposefile, err := os.ReadFile("./docker-compose.yml")
		k1, err := loader.ParseYAML(dockercomposefile)
		tmpk := k1["services"]
		tmp3, _ := tmpk.(map[string]interface{})
		services, err := LoadServices(tmp3)
		containers, err := TransServicesToContainer(services)
		var f1  struct {
		 	Containers Containers
		}
		f1.Containers = containers
		batectyaml, err := yaml.Marshal(&f1)
	  if err != nil {
	  	fmt.Printf("error: %v", err)
	  }
		os.WriteFile("./batect.yml", batectyaml,0666)
}
