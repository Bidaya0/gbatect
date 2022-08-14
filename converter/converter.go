package converter

import (
	"fmt"
	"github.com/compose-spec/compose-go/loader"
	composetypes "github.com/compose-spec/compose-go/types"
	batecttypes "github.com/bidaya0/gbatect/types"
//	"gopkg.in/yaml.v3"
//	"os"
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
			Environment: service.Environment,
		}
		if service.Build != nil {
			containeroption.BuildDirectory = service.Build.Context
			containeroption.Dockerfile = service.Build.Dockerfile
		}
		if len(service.Entrypoint) > 0 {
			containeroption.Entrypoint = strings.Join(service.Entrypoint, " ")
		}
		for _, port := range service.Ports {
			portstring := fmt.Sprintf("%v:%v", port.Published, port.Target)
			containeroption.Ports = append(containeroption.Ports, portstring)
		}
		for _, volume := range service.Volumes {
			volumestring := fmt.Sprintf("%v:%v", volume.Source, volume.Target)
			containeroption.Volumes = append(containeroption.Volumes, volumestring)
		}

		containers[service.Name] = containeroption
	}
	return containers, nil
}

