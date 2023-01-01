package main

import (
	"errors"
	"os"

	"github.com/Abeldlp/gortisan/cmd"
)

func checkFileExists() bool {
	_, error := os.Stat("go.mod")
	return !errors.Is(error, os.ErrNotExist)
}

func main() {
	if !checkFileExists() {
		panic("You must be in a golang project to use gortisan")
	}
	cmd.RootCmd.AddCommand(cmd.MakeCmd, cmd.ConfigCmd)
	cmd.RootCmd.Execute()
}
