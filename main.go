package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmdInfrastructureCharts := createInfrastructureChartsCMD()
	var rootCmd = &cobra.Command{
		Use: "iits-chart-creator",
	}
	rootCmd.AddCommand(cmdInfrastructureCharts)
	cobra.CheckErr(rootCmd.Execute())
}

func createInfrastructureChartsCMD() *cobra.Command {
	var cmdCreateIngress = &cobra.Command{
		Use:   "ingress",
		Short: "Creates a Ingress YAML deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, "ingress.yaml", "infrastructure-charts")
		},
	}

	var cmdCreateHelpers = &cobra.Command{
		Use:   "helpers.tpl",
		Short: "Creates a helpers.tpl deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, "_helpers.tpl", "infrastructure-charts")
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
			checkAndCreateResource(cmd, "deployment.yaml", "infrastructure-charts")
		},
	}

	var serviceCreateHelpers = &cobra.Command{
		Use:   "service",
		Short: "Creates a service deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, "service.yaml", "infrastructure-charts")
		},
	}

	var saCreateHelpers = &cobra.Command{
		Use:   "serviceaccount",
		Short: "Creates a serviceaccount deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, "serviceaccount.yaml", "infrastructure-charts")
		},
	}

	var serviceMonitorCreateHelpers = &cobra.Command{
		Use:   "service-monitor",
		Short: "Creates a service-monitor deployment file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, "service-monitor.yaml", "infrastructure-charts")
		},
	}

	cmdInfrastructureCharts.AddCommand(cmdCreateIngress)
	cmdInfrastructureCharts.AddCommand(cmdCreateHelpers)
	cmdInfrastructureCharts.AddCommand(saCreateHelpers)
	cmdInfrastructureCharts.AddCommand(serviceCreateHelpers)
	cmdInfrastructureCharts.AddCommand(deploymentCreateHelpers)
	cmdInfrastructureCharts.AddCommand(serviceMonitorCreateHelpers)
	return cmdInfrastructureCharts
}
