package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type CliBuildpackUsage struct{}

func (c *CliBuildpackUsage) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "CliBuildpackUsage",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 1,
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

func (c *CliBuildpackUsage) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("----------              FIN               -----------")
}
