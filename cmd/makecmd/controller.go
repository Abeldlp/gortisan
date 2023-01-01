package makecmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var FileExtension = ".go"

func formatFilename(name string) string {
	var mySli []string

	for _, val := range strings.Split(name, "-") {
		mySli = append(mySli, strings.Title(val))
	}

	return strings.Join(mySli, "")
}

func writeControllerBoilerplate(name string) string {
	controllerBoilerPlate := `
package controllers

import "github.com/gin-gonic/gin"

func GetAll` + formatFilename(name) + `(*gin.Context) {
	//TODO: implement GET
}

func Get` + formatFilename(name) + `ById(*gin.Context) {
	//TODO: implement GET
}

func Create` + formatFilename(name) + `(*gin.Context) {
	//TODO: implement GET
}

func Update` + formatFilename(name) + `(*gin.Context) {
	//TODO: implement GET
}
    
func Delete` + formatFilename(name) + `(*gin.Context) {
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
		_, err := os.Stat("controllers")

		if os.IsNotExist(err) {
			errDir := os.Mkdir("controllers", 0755)
			if errDir != nil {
				log.Fatal(err)
			}
		}

		fileName := strings.Join(args, " ")

		f, err := os.Create(filepath.Join("controllers", filepath.Base(fileName+FileExtension)))
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = f.WriteString(writeControllerBoilerplate(fileName))
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

		fmt.Printf("%v controller created successfully\n", fileName)
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
