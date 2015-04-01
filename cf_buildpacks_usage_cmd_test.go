package main_test

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/csterwa/cf_buildpacks_usage_cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cloud Foundry Buildpack Usage Command", func() {
	Describe(".Run", func() {
		var fakeCliConnection *fakes.FakeCliConnection
		var callBuildpackUsageCommandPlugin *CliBuildpackUsage

		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			callBuildpackUsageCommandPlugin = &CliBuildpackUsage{}
		})

		It("calls the buildpack-usage command with no arguments", func() {
			fakeAppsResponse := []string{"{\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Node.js\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
			output := io_helpers.CaptureOutput(func() {
				callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage"})
			})

			Expect(output[1]).To(Equal("Buildpacks in use across all organizations..."))
			Expect(output[2]).To(Equal("2 apps found"))
			Expect(output[4]).To(ContainSubstring("App Name"))
			Expect(output[4]).To(ContainSubstring("Buildpack"))
			Expect(output[4]).To(ContainSubstring("Detected Buildpack"))

			dataRows := []string{output[5], output[6]}
			dataRowsAsS := strings.Join(dataRows, "")

			fmt.Println(output[5])
			fmt.Println(output[6])
			Expect(dataRowsAsS).To(ContainSubstring("app2"))
			Expect(dataRowsAsS).To(ContainSubstring("Java"))
			Expect(dataRowsAsS).To(ContainSubstring("app1"))
			Expect(dataRowsAsS).To(ContainSubstring("Node.js"))
		})
	})
})
