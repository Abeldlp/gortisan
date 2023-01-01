package cmd

import (
	"github.com/Abeldlp/gortisan/cmd/makecmd"
	"github.com/spf13/cobra"
)

var MakeCmd = &cobra.Command{
	Use:   "make [resource]",
	Short: "Make specified resource",
	Long: `
The make command creates a resource based on the name passed.
The name must be passed with a dash and lower letters.

With the make command you can make a controller, a model or a route.

Eg: hello-world
	`,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	MakeCmd.AddCommand(makecmd.ControllerCmd, makecmd.ModelCmd, makecmd.RouteCmd)
}
