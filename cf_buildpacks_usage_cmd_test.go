package main_test

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/mkuratczyk/cf_buildpacks_usage_cmd"
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
			fakeAppsResponse := []string{"{\"total_pages\":1,\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Node.js\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
			output := io_helpers.CaptureOutput(func() {
				callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage"})
			})

			Expect(output[1]).To(Equal("2 buildpacks found across 2 app deployments"))
			Expect(output[3]).To(Equal("Buildpacks Used"))
			Expect(output[5]).To(Equal(fmt.Sprintf("%-8v%-24v", "Count", "Name")))
			Expect(output[7]).To(ContainSubstring("Java"))
			Expect(output[8]).To(ContainSubstring("Node.js"))
		})

		It("removes duplicates from buildpacks used list", func() {
			fakeAppsResponse := []string{"{\"total_pages\":1,\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Java\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
			output := io_helpers.CaptureOutput(func() {
				callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage"})
			})

			Expect(output[1]).To(Equal("1 buildpacks found across 2 app deployments"))
			Expect(output[3]).To(Equal("Buildpacks Used"))
			Expect(output[7]).To(ContainSubstring("Java"))
		})

		It("counts the amount of each buildpack used", func() {
			fakeAppsResponse := []string{"{\"total_pages\":1,\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Java\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
			output := io_helpers.CaptureOutput(func() {
				callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage"})
			})

			Expect(output[7]).To(Equal(fmt.Sprintf("%-8v%-24v", "2", "Java")))
		})

		It("pages through all app data to combine results", func() {
			fakeAppsResponse := []string{"{\"total_pages\":2,\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Java\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
			output := io_helpers.CaptureOutput(func() {
				callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage"})
			})

			Expect(output[7]).To(Equal(fmt.Sprintf("%-8v%-24v", "4", "Java")))
		})

		Context("when called using the --apps flag", func() {
			It("prints which apps use a given buildpack", func() {
				fakeAppsResponse := []string{"{\"total_pages\":1,\"total_results\":2,\"resources\":[{\"entity\":{\"name\":\"app1\",\"buildpack\":null,\"detected_buildpack\":\"Java\"}},{\"entity\":{\"name\":\"app2\",\"buildpack\":\"Java\",\"detected_buildpack\":null}}]}"}
				fakeCliConnection.CliCommandWithoutTerminalOutputReturns(fakeAppsResponse, nil)
				output := io_helpers.CaptureOutput(func() {
					callBuildpackUsageCommandPlugin.Run(fakeCliConnection, []string{"buildpack-usage", "--apps"})
				})

				Expect(output[7]).To(Equal(fmt.Sprintf("%-8v%-24v%v", "2", "Java", "app1, app2")))
			})
		})
	})
})
