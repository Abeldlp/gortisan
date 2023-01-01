package makecmd

import (
	"strings"

	"github.com/spf13/cobra"
)

func writeModelBoilerplate(name string) string {
	modelBoilerPlate := `
package models 

type ` + FormatFilename(name) + ` struct {
    //TODO: implement
}

func New` + FormatFilename(name) + `() ` + FormatFilename(name) + ` {
    return ` + FormatFilename(name) + `{}
}
`
	return modelBoilerPlate
}

var ModelCmd = &cobra.Command{
	Use:   "model [name of the controller]",
	Short: "Creates a model",
	Long: `
Creates a model based on the name passed in the models folder.
The name must be passed with a dash and lower letters.

Eg: hello-world
    `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := strings.Join(args, " ")
		WriteFile("model", writeModelBoilerplate(fileName), args)
	},
}
