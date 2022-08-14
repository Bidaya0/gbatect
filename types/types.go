package types

import (
//	"fmt"
//	"github.com/compose-spec/compose-go/loader"
	composetypes "github.com/compose-spec/compose-go/types"
//	"gopkg.in/yaml.v3"
//	"os"
//	"strings"
)


type ContainerOption struct {
	Image          string                  ",omitempty"
	BuildDirectory string                  "build_directory,omitempty"
	Dockerfile     string                  ",omitempty"
	Entrypoint     string                  ",omitempty"
	Ports          []string                ",omitempty"
	Volumes        []string                ",omitempty"
	Environment    composetypes.MappingWithEquals ",omitempty"
	BuildArgs      map[string]string       "build_args,omitempty"
}

type ContainerItem struct {
	Name    string
	Options map[string]ContainerOption
}

type Containers map[string]ContainerOption

type BatectConfig struct {
	Containers Containers
}
