package cmd

import (
	"github.com/Abeldlp/gortisan/cmd/makecmd"
	"github.com/spf13/cobra"
)

var MakeCmd = &cobra.Command{
	Use:   "make [string to print]",
	Short: "make specified resource",
	Long:  "something is being made out of this command",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	MakeCmd.AddCommand(makecmd.ControllerCmd, makecmd.ModelCmd)
}
