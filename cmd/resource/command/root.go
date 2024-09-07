package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ai-resource-exporter",
	Short: "ai-resource-exporter is a hardware resource collector on the host node.",
	Long: `ai-resource-exporter is the cpu, numa, network, gpu topology, cadvisor
container monitoring and other information on the host node.`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	rootCmd.Execute()
}
