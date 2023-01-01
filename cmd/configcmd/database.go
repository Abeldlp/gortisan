package configcmd

import "github.com/spf13/cobra"

func writeDatabaseConfigBoilerplate() string {
	configBoilerplate := `
package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate() // Add models here

	DB = db

	return db
}
	`

	return configBoilerplate
}

var DatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Configures database",
	Long: `
Configures database based on the name passed in the config folder.

		`,
	Run: func(cmd *cobra.Command, args []string) {
		WriteFile("database", writeDatabaseConfigBoilerplate())
	},
}
