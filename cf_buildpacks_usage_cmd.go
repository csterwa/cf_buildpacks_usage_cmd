package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

type CliBuildpackUsage struct{}

type AppSearchResults struct {
	TotalResults int                  `json:"total_results"`
	Resources    []AppSearchResources `json:"resources"`
}

type AppSearchResources struct {
	Entity AppSearchEntity `json:"entity"`
}

type AppSearchEntity struct {
	Name              string `json:"name"`
	Buildpack         string `json:"buildpack"`
	DetectedBuildpack string `json:"detected_buildpack"`
}

func (c *CliBuildpackUsage) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "CliBuildpackUsage",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 2,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "buildpack-usage",
				HelpText: "Command to view buildpack usage at org, space or entire Cloud Foundry levels.",
				UsageDetails: plugin.Usage{
					Usage: "buildpack-usage\n   cf buildpack-usage [-o ORG] [-s SPACE]",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(CliBuildpackUsage))
}

func (c CliBuildpackUsage) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("")
	fmt.Println("Buildpacks in use across all organizations...")

	res := c.GetAppData(cliConnection)

	fmt.Printf("| App Name | Buildpack | Detected Buildpack |\n")
	for _, val := range res {
		fmt.Printf("| %v | %v | %v |\n", val.Entity.Name, val.Entity.Buildpack, val.Entity.DetectedBuildpack)
	}
}

func (c CliBuildpackUsage) GetAppData(cliConnection plugin.CliConnection) []AppSearchResources {
	cmd := []string{"curl", "/v2/apps"}
	output, _ := cliConnection.CliCommandWithoutTerminalOutput(cmd...)
	res := &AppSearchResults{}
	json.Unmarshal([]byte(strings.Join(output, "")), &res)
	fmt.Printf("%v apps found\n\n", res.TotalResults)

	return res.Resources
}
