package manager

import (
	"fmt"

	"github.com/spf13/cobra"
	"semay.com/admin/database"
	"semay.com/admin/models"
)

var (
	testCmd = &cobra.Command{
		Use:   "testcmd",
		Short: "Seed Intial data to the database ",
		Long:  `It updates databased from provided file in the root directory from where the app is running. The CSV file should be in the provided sample format on the Repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			testcmd()
		},
	}
)

func testcmd() {

	fmt.Println("CMD test bed")
	db := database.ReturnSession()
	roles := []models.Role{}
	user := models.User{}
	app := models.App{}
	db.Model(&models.User{}).Where("id = ?", 1).Find(&user)
	db.Model(&models.Role{}).Order("id asc").Find(&roles)
	db.Model(&models.App{}).Where("id = ?", 1).Find(&app)

	fmt.Println(roles)
	fmt.Println(user)
	fmt.Println(app)
	for i := range roles {
		tx := db.Begin()
		db.Model(&user).Association("Roles").Append(&roles[i])
		db.Model(&app).Association("Roles").Append(&roles[i])
		tx.Commit()
	}

	// db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(roles, len(roles))
	// db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(users, len(users))
	// db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(apps, len(apps))
	// db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(features, len(features))
	// db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(pages, len(pages))

}

func init() {
	goBlueCmd.AddCommand(testCmd)

}
