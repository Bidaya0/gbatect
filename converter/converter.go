package converter

import (
	"fmt"
	"github.com/compose-spec/compose-go/loader"
	composetypes "github.com/compose-spec/compose-go/types"
	batecttypes "github.com/bidaya0/gbatect/types"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func LoadServices(servicesDict map[string]interface{}) ([]composetypes.ServiceConfig, error) {
	var services []composetypes.ServiceConfig
	for name, element := range servicesDict {
		serviceConfig := &composetypes.ServiceConfig{
			Name:  name,
			Scale: 1,
		}
		loader.Transform(element, serviceConfig)
		services = append(services, *serviceConfig)
	}
	return services, nil
}

func TransServicesToContainer(servicesconfigs []composetypes.ServiceConfig) (batecttypes.Containers, error) {
	containers := make(map[string]batecttypes.ContainerOption)
	for _, service := range servicesconfigs {
		containeroption := batecttypes.ContainerOption{
			Image:       service.Image,
			AdditionalHosts: service.ExtraHosts,
			CapabilitiesToAdd: service.CapAdd,
			CapabilitiesToDrop: service.CapDrop,
			Devices: service.Devices,
			EnableInitProcess: service.Init,
			Environment: service.Environment,
			ShmSize: service.ShmSize,
			WorkingDirectory: service.WorkingDir,
		}
		for _, network := range service.Networks {
		  containeroption.AdditionalHostnames = append(containeroption.AdditionalHostnames, network.Aliases...)
		}
		if service.Build != nil {
			containeroption.BuildArgs = service.Build.Args
			containeroption.BuildDirectory = service.Build.Context
			containeroption.Dockerfile = service.Build.Dockerfile
			containeroption.BuildTarget = service.Build.Target
		}
		if service.HealthCheck != nil {
			containeroption.HealthCheck = batecttypes.HealthCheck{
				Interval: service.HealthCheck.Interval,	
				Retries: service.HealthCheck.Retries,	
				StartPeriod: service.HealthCheck.StartPeriod,	
				Timeout: service.HealthCheck.Timeout,	
			}
			if service.HealthCheck.Test != nil {
				containeroption.HealthCheck.Command = strings.Join(service.HealthCheck.Test, " ")
			}
		}
		if service.Logging != nil {
			containeroption.LogDriver = service.Logging.Driver
			containeroption.LogOptions = service.Logging.Options
		}
		if len(service.Entrypoint) > 0 {
			containeroption.Entrypoint = strings.Join(service.Entrypoint, " ")
		}
		if len(service.Command) > 0 {
			containeroption.Command = strings.Join(service.Command, " ")
		}
		for _, port := range service.Ports {
			portstring := fmt.Sprintf("%v:%v", port.Published, port.Target)
			containeroption.Ports = append(containeroption.Ports, portstring)
		}
		for _, volume := range service.Volumes {
			volumestring := fmt.Sprintf("%v:%v", volume.Source, volume.Target)
			containeroption.Volumes = append(containeroption.Volumes, volumestring)
		}
		for service_name, _ := range service.DependsOn {
			containeroption.Dependencies = append(containeroption.Dependencies, service_name)
		}

		containers[service.Name] = containeroption
	}
	return containers, nil
}

func ReadAndConvert(sourceFilePath string) {
	dockercomposefile, err := os.ReadFile(sourceFilePath)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	k1, err := loader.ParseYAML(dockercomposefile)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	tmpk := k1["services"]
	tmp3, _ := tmpk.(map[string]interface{})
	services, err := LoadServices(tmp3)
	containers, err := TransServicesToContainer(services)
	var f1 batecttypes.BatectConfig 
	f1.Containers = containers
	batectyaml, err := yaml.Marshal(&f1)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf(string(batectyaml))
}

