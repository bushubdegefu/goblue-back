package manager

import (
	"github.com/spf13/cobra"
	"semay.com/admin/models"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate Updated Models to the Database",
		Long:  `Migrate Updated Models to the Database. The database URI is to be provided within the migrate function or as .env variable`,
		Run: func(cmd *cobra.Command, args []string) {
			migrate()
		},
	}
)

func migrate() {

	models.InitDatabase()
}

func init() {
	goBlueCmd.AddCommand(migrateCmd)

}
