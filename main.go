package main

import "github.com/Abeldlp/gortisan/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.MakeCmd)
	cmd.RootCmd.Execute()
}
