package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

// CliBuildpackUsage represents Buildpack Usage CLI interface
type CliBuildpackUsage struct{}

// AppSearchResults represents top level attributes of JSON response from Cloud Foundry API
type AppSearchResults struct {
	TotalResults int                  `json:"total_results"`
	TotalPages   int                  `json:"total_pages"`
	Resources    []AppSearchResources `json:"resources"`
}

// AppSearchResources represents resources attribute of JSON response from Cloud Foundry API
type AppSearchResources struct {
	Entity AppSearchEntity `json:"entity"`
}

// AppSearchEntity represents entity attribute of resources attribute within JSON response from Cloud Foundry API
type AppSearchEntity struct {
	Name              string `json:"name"`
	Buildpack         string `json:"buildpack"`
	DetectedBuildpack string `json:"detected_buildpack"`
}

// GetMetadata provides the Cloud Foundry CLI with metadata to provide user about how to use buildpack-usage command
func (c *CliBuildpackUsage) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "CliBuildpackUsage",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "buildpack-usage",
				HelpText: "Command to view buildpack usage in current CLI target context.",
				UsageDetails: plugin.Usage{
					Usage: "buildpack-usage\n   cf buildpack-usage",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(CliBuildpackUsage))
}

// Run is what is executed by the Cloud Foundry CLI when the buildpack-usage command is specified
func (c CliBuildpackUsage) Run(cliConnection plugin.CliConnection, args []string) {
	res := c.GetAppData(cliConnection)

	var buildpacksUsed sort.StringSlice

	for _, val := range res.Resources {
		bp := val.Entity.Buildpack
		if bp == "" {
			bp = val.Entity.DetectedBuildpack
		}
		buildpacksUsed = append(buildpacksUsed, bp)
	}

	var buildpackUsageTable = c.CreateBuildpackUsageTable(buildpacksUsed)
	c.PrintBuildpacks(buildpackUsageTable, res.TotalResults)
}

// CreateBuildpackUsageTable creates a map whose key is buildpack and value is count of that buildpack
func (c CliBuildpackUsage) CreateBuildpackUsageTable(buildpacksUsed sort.StringSlice) map[string]int {
	buildpackUsageCounts := make(map[string]int)

	for _, buildpackName := range buildpacksUsed {
		if _, ok := buildpackUsageCounts[buildpackName]; ok {
			buildpackUsageCounts[buildpackName]++
		} else {
			buildpackUsageCounts[buildpackName] = 1
		}
	}

	return buildpackUsageCounts
}

// PrintBuildpacks prints the buildpack data to console
func (c CliBuildpackUsage) PrintBuildpacks(buildpackUsageTable map[string]int, totalResults int) {
	fmt.Println("")
	fmt.Printf("%v buildpacks found across %v app deployments\n\n", len(buildpackUsageTable), totalResults)
	fmt.Println("Buildpacks Used\n")
	fmt.Println("Count\tName")
	fmt.Println("-------------------------------")

	buildpackNames := make([]string, len(buildpackUsageTable))
	j := 0
	for buildpack, _ := range buildpackUsageTable {
		buildpackNames[j] = buildpack
		j++
	}

	sort.Strings(buildpackNames)
	for _, buildpack := range buildpackNames {
		fmt.Printf("%v\t%v\n", buildpackUsageTable[buildpack], buildpack)
	}
}

// GetAppData requests all of the Application data from Cloud Foundry
func (c CliBuildpackUsage) GetAppData(cliConnection plugin.CliConnection) AppSearchResults {
	var res AppSearchResults
	res = c.UnmarshallAppSearchResults("/v2/apps?order-direction=asc&results-per-page=100", cliConnection)

	if res.TotalPages > 1 {
		for i := 2; i <= res.TotalPages; i++ {
			apiUrl := fmt.Sprintf("/v2/apps?order-direction=asc&page=%v&results-per-page=100", strconv.Itoa(i))
			tRes := c.UnmarshallAppSearchResults(apiUrl, cliConnection)
			res.Resources = append(res.Resources, tRes.Resources...)
		}
	}

	return res
}

func (c CliBuildpackUsage) UnmarshallAppSearchResults(apiUrl string, cliConnection plugin.CliConnection) AppSearchResults {
	var tRes AppSearchResults
	cmd := []string{"curl", apiUrl}
	output, _ := cliConnection.CliCommandWithoutTerminalOutput(cmd...)
	json.Unmarshal([]byte(strings.Join(output, "")), &tRes)

	return tRes
}
