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
	AdditionalHostnames      []string       "additional_hostnames,omitempty"
	AdditionalHosts      map[string]string       "additional_hosts,omitempty"
	BuildArgs      map[string]string       "build_args,omitempty"
	BuildDirectory string                  "build_directory,omitempty"
	BuildStage      string       "build_stage,omitempty"
	BuildTarget      string       "build_target,omitempty"
	CapabilitiesToAdd      []string       "capabilities_to_add,omitempty"
	CapabilitiesToDrop      []string       "capabilities_to_drop,omitempty"
	Command     string                  "command,omitempty"
	Dependencies     []string                  "dependencies,omitempty"
	Devices     []string                  "devices,omitempty" // /dev/sda:/dev/disk:r ??
	Dockerfile     string                  ",omitempty"	
	EnableInitProcess bool                "enable_init_process,omitempty"
	Entrypoint     string                  ",omitempty"
	Environment    composetypes.MappingWithEquals ",omitempty"
	HealthCheck			HealthCheck                  "health_check,omitempty"
	Image          string                  ",omitempty"
	ImagePullPolicy     string                  "image_pull_policy,omitempty"
	LogDriver     string                  "log_driver,omitempty"
	LogOptions     map[string]string                  "log_options,omitempty"
	Ports          []string                ",omitempty"
	Privileged    bool                  "privileged,omitempty"
	RunAsCurrentUser   RunAsCurrentUser                  "run_as_current_user,omitempty"
	setup_commands   []SetupCommand                  "setup_commands,omitempty"
	ShmZize   string                  "shm_size,omitempty"
	Volumes        []string                ",omitempty"
	WorkingDirectory   string                  "working_directory,omitempty"
}

type SetupCommand struct{
	Command   string                  "command,omitempty"
	workingDirectory   string                  "working_directory,omitempty"
}

type RunAsCurrentUser struct{
	Enabled   bool                  "enabled,omitempty"
	HomeDirectory   bool                  "home_directory,omitempty"
}

type HealthCheck struct {
	Command     string                  ",omitempty"
	Retries     string                  ",omitempty"
	Interval     string                  ",omitempty"
	StartPeriod     string                  "start_period,omitempty"
	Timeout    string                  "timeout,omitempty"
}

type ContainerItem struct {
	Name    string
	Options map[string]ContainerOption
}

type Containers map[string]ContainerOption

type BatectConfig struct {
	Containers Containers
}
