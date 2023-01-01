package cmd

import (
	"github.com/Abeldlp/gortisan/cmd/configcmd"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config [resource]",
	Short: "Configures resource",
	Long: `
The config command configures a resource based on the name passed.

Available resources: database, docker

Eg: database
	`,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	ConfigCmd.AddCommand(configcmd.DatabaseCmd)
}
