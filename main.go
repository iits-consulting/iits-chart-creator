package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var cmdCreateIngress = &cobra.Command{
		Use:   "ingress",
		Short: "Creates a Ingress YAML file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, "ingress.yaml")
		},
	}

	var cmdCreateHelpers = &cobra.Command{
		Use:   "helpers.tpl",
		Short: "Creates a helpers.tpl file",
		Run: func(cmd *cobra.Command, args []string) {
			checkAndCreateResource(cmd, "_helpers.tpl")
		},
	}

	var cmdInfrastructureCharts = &cobra.Command{
		Use:   "infrastructure-charts",
		Short: "Creates Infrastructure Charts",
	}

	cmdInfrastructureCharts.AddCommand(cmdCreateIngress)
	cmdInfrastructureCharts.AddCommand(cmdCreateHelpers)

	var rootCmd = &cobra.Command{Use: "helm iits-chart-creator"}
	rootCmd.AddCommand(cmdInfrastructureCharts)
	cobra.CheckErr(rootCmd.Execute())
}
