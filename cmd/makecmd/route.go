package makecmd

import (
	"strings"

	"github.com/spf13/cobra"
)

func writeRouteBouterplate(name string) string {
	routeBoilerPlate := `
package routes

import (
	"github.com/gin-gonic/gin"
)

func Initialize` + FormatFilename(name) + `Routes(router *gin.Engine) {
	router = router.Group("/` + strings.ToLower(name) + `")

	{
		router.GET("/", controllers.GetAll` + FormatFilename(name) + `)
		router.GET("/:id", controllers.Get` + FormatFilename(name) + `ById)
		router.POST("/", controllers.Create` + FormatFilename(name) + `)
		router.PUT("/:id", controllers.Update` + FormatFilename(name) + `)
		router.DELETE("/:id", controllers.Delete` + FormatFilename(name) + `)
	}
}
`

	return routeBoilerPlate
}

var RouteCmd = &cobra.Command{
	Use:   "route [name of the route]",
	Short: "Creates a route",
	Long: `
Creates a route based on the name passed in the routes folder.
The name must be passed with a dash and lower letters.

Eg: hello-world
		`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := strings.Join(args, " ")
		WriteFile("route", writeRouteBouterplate(fileName), args)
	},
}
