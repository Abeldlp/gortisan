package makecmd

import (
	"strings"

	"github.com/spf13/cobra"
)

func writeControllerBoilerplate(name string) string {
	controllerBoilerPlate := `
package controllers

import "github.com/gin-gonic/gin"

func GetAll` + FormatFilename(name) + `(*gin.Context) {
	//TODO: implement GET
}

func Get` + FormatFilename(name) + `ById(*gin.Context) {
	//TODO: implement GET
}

func Create` + FormatFilename(name) + `(*gin.Context) {
	//TODO: implement GET
}

func Update` + FormatFilename(name) + `(*gin.Context) {
	//TODO: implement GET
}
    
func Delete` + FormatFilename(name) + `(*gin.Context) {
	//TODO: implement GET
}
`
	return controllerBoilerPlate
}

var ControllerCmd = &cobra.Command{
	Use:   "controller [name of the controller]",
	Short: "Creates a controller",
	Long: `
Creates a controller based on the name passed in the controllers folder.
The name must be passed with a dash and lower letters.

Eg: hello-world
    `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := strings.Join(args, " ")
		WriteFile("controller", writeControllerBoilerplate(fileName), args)
	},
}
