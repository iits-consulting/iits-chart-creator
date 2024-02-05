package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "dev"
var overrideReleaseName string

func main() {
	cmdInfrastructureCharts := createInfrastructureChartsCMD()
	var rootCmd = &cobra.Command{
		Use:     "iits-chart-creator",
		Version: version,
	}
	folderName, err := getCurrentFolderName()
	if err != nil {
		fmt.Printf("Error getting current folder name: %v\n", err)
		return
	}
	rootCmd.Version = version
	rootCmd.PersistentFlags().StringVarP(&overrideReleaseName, "override-release-name", "o", folderName, "Override the chart's release name")
	rootCmd.AddCommand(cmdInfrastructureCharts)
	cobra.CheckErr(rootCmd.Execute())
}

func createInfrastructureChartsCMD() *cobra.Command {
	var cmdCreateIngress = &cobra.Command{
		Use:   "ingress",
		Short: "Creates a Ingress YAML deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, DataInputForCreation{
				ManifestName: "ingress.yaml",
				ChartType:    "infrastructure-charts",
				Version:      version,
				ReleaseName:  overrideReleaseName,
			})
		},
	}

	var cmdCreateHelpers = &cobra.Command{
		Use:   "helpers.tpl",
		Short: "Creates a helpers.tpl deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, DataInputForCreation{
				ManifestName: "_helpers.tpl",
				ChartType:    "infrastructure-charts",
				Version:      version,
				ReleaseName:  overrideReleaseName,
			})
		},
	}

	var cmdInfrastructureCharts = &cobra.Command{
		Use:   "infrastructure-charts",
		Short: "Creates Infrastructure Charts",
	}

	var deploymentCreateHelpers = &cobra.Command{
		Use:   "deployment",
		Short: "Creates a deployment deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, DataInputForCreation{
				ManifestName: "deployment.yaml",
				ChartType:    "infrastructure-charts",
				Version:      version,
				ReleaseName:  overrideReleaseName,
			})
		},
	}

	var serviceCreateHelpers = &cobra.Command{
		Use:   "service",
		Short: "Creates a service deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, DataInputForCreation{
				ManifestName: "service.yaml",
				ChartType:    "infrastructure-charts",
				Version:      version,
				ReleaseName:  overrideReleaseName,
			})
		},
	}

	var saCreateHelpers = &cobra.Command{
		Use:   "serviceaccount",
		Short: "Creates a serviceaccount deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, DataInputForCreation{
				ManifestName: "serviceaccount.yaml",
				ChartType:    "infrastructure-charts",
				Version:      version,
				ReleaseName:  overrideReleaseName,
			})
		},
	}

	var serviceMonitorCreateHelpers = &cobra.Command{
		Use:   "service-monitor",
		Short: "Creates a service-monitor deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, DataInputForCreation{
				ManifestName: "service-monitor.yaml",
				ChartType:    "infrastructure-charts",
				Version:      version,
				ReleaseName:  overrideReleaseName,
			})
		},
	}

	var podMonitorCreateHelpers = &cobra.Command{
		Use:   "pod-monitor",
		Short: "Creates a pod-monitor deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, DataInputForCreation{
				ManifestName: "pod-monitor.yaml",
				ChartType:    "infrastructure-charts",
				Version:      version,
				ReleaseName:  overrideReleaseName,
			})
		},
	}

	cmdInfrastructureCharts.AddCommand(cmdCreateIngress)
	cmdInfrastructureCharts.AddCommand(cmdCreateHelpers)
	cmdInfrastructureCharts.AddCommand(saCreateHelpers)
	cmdInfrastructureCharts.AddCommand(serviceCreateHelpers)
	cmdInfrastructureCharts.AddCommand(deploymentCreateHelpers)
	cmdInfrastructureCharts.AddCommand(serviceMonitorCreateHelpers)
	cmdInfrastructureCharts.AddCommand(podMonitorCreateHelpers)
	return cmdInfrastructureCharts
}
