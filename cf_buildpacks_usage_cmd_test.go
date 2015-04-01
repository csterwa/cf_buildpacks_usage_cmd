package main_test

import (
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
			Expect(output[2]).To(Equal("2 buildpacks found across 2 app deployments"))
			Expect(output[4]).To(Equal("Buildpacks Used"))
			Expect(output[5]).To(Equal("Java"))
			Expect(output[6]).To(Equal("Node.js"))
		})

		It("removes duplicates from buildpacks used list", func() {
			fakeAppsResponse := []string{"{\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Java\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
			output := io_helpers.CaptureOutput(func() {
				callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage"})
			})

			Expect(output[1]).To(Equal("Buildpacks in use across all organizations..."))
			Expect(output[2]).To(Equal("1 buildpacks found across 2 app deployments"))
			Expect(output[4]).To(Equal("Buildpacks Used"))
			Expect(output[5]).To(Equal("Java"))
		})
	})
})
