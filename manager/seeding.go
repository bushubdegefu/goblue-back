package manager

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
)

var (
	seedCmd = &cobra.Command{
		Use:   "seed",
		Short: "Seed Intial data to the database ",
		Long:  `It updates databased from provided file in the root directory from where the app is running. The CSV file should be in the provided sample format on the Repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			seeding()
		},
	}
)

func seeding() {
	// getting database session
	db := database.ReturnSession()
	// definning list of models to be populated first
	users := make([]models.User, 0)
	roles := make([]models.Role, 0)
	apps := make([]models.App, 0)
	features := make([]models.Feature, 0)
	pages := make([]models.Page, 0)

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open("seeding.csv")

	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	// Closes the file
	defer file.Close()

	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	// Loop to iterate through
	// and print each of the string slice
	for _, eachrecord := range records {
		switch eachrecord[0] {
		case "roles":
			single_map := map[string]interface{}{}
			single_struct := models.Role{}
			for _, x := range eachrecord[1:] {
				values := strings.Split(x, ":")
				if len(values) > 1 {

					fieldName := values[0]
					fieldValue := values[1]
					single_map[fieldName] = fieldValue

					mapstructure.Decode(single_map, &single_struct)
				}
			}
			roles = append(roles, single_struct)
			tx := db.Begin()
			db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(roles, len(roles))
			tx.Commit()
		case "users":
			single_map := map[string]interface{}{}
			single_struct := models.User{}
			for _, x := range eachrecord[1:] {
				values := strings.Split(x, ":")
				if len(values) > 1 {
					if values[0] != "disabled" {
						fieldName := values[0]
						fieldValue := values[1]
						single_map[fieldName] = fieldValue

					} else {
						fieldName := values[0]
						fieldValue, _ := strconv.ParseBool(values[1])
						single_map[fieldName] = fieldValue
					}

					mapstructure.Decode(single_map, &single_struct)
				}

			}
			users = append(users, single_struct)
			tx := db.Begin()
			db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(users, len(users))
			tx.Commit()
		case "features":
			single_map := map[string]interface{}{}
			single_struct := models.Feature{}
			for _, x := range eachrecord[1:] {
				values := strings.Split(x, ":")
				if len(values) > 1 {

					fieldName := values[0]
					fieldValue := values[1]
					single_map[fieldName] = fieldValue

					mapstructure.Decode(single_map, &single_struct)
				}

			}
			features = append(features, single_struct)
			tx := db.Begin()
			db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(features, len(features))
			tx.Commit()
		case "app":
			single_map := map[string]interface{}{}
			single_struct := models.App{}
			for _, x := range eachrecord[1:] {
				values := strings.Split(x, ":")
				if len(values) > 1 {

					fieldName := values[0]
					fieldValue := values[1]
					single_map[fieldName] = fieldValue

					mapstructure.Decode(single_map, &single_struct)
				}

			}
			apps = append(apps, single_struct)
			tx := db.Begin()
			db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(apps, len(apps))
			tx.Commit()
		case "pages":
			single_map := map[string]interface{}{}
			single_struct := models.Page{}
			for _, x := range eachrecord[1:] {
				values := strings.Split(x, ":")
				if len(values) > 1 {

					fieldName := values[0]
					fieldValue := values[1]
					single_map[fieldName] = fieldValue

					mapstructure.Decode(single_map, &single_struct)
				}

			}
			pages = append(pages, single_struct)
			tx := db.Begin()
			db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(pages, len(pages))
			tx.Commit()
		case "rolefeatures":
			feature := models.Feature{}
			role := models.Role{}
			for _, x := range eachrecord[1:] {
				values := strings.Split(x, ":")
				if len(values) > 1 {
					tx := db.Begin()
					db.Model(&models.Role{}).Where("name = ?", values[0]).Find(&role)
					db.Model(&models.Feature{}).Where("name = ?", values[1]).Find(&feature)
					db.Model(&role).Association("Features").Append(&feature)
					tx.Commit()
				}
			}
		case "approles":
			app := models.App{}
			role := models.Role{}
			for _, x := range eachrecord[1:] {
				values := strings.Split(x, ":")
				if len(values) > 1 {
					tx := db.Begin()
					db.Model(&models.Role{}).Where("name = ?", values[1]).Find(&role)
					db.Model(&models.App{}).Where("name = ?", values[0]).Find(&app)
					db.Model(&app).Association("Roles").Append(&role)
					tx.Commit()
				}
			}
		case "pagefeatures":
			// feature := models.Feature{}
			// page := models.Page{}
			// for _, x := range eachrecord[1:] {
			// 	values := strings.Split(x, ":")
			// 	if len(values) > 1 {
			// 		tx := db.Begin()
			// 		db.Model(&models.Feature{}).Where("name = ?", values[1]).Find(&feature)
			// 		db.Model(&models.Page{}).Where("name = ?", values[0]).Find(&page)
			// 		db.Model(&page).Association("Features").Append(&feature)
			// 		tx.Commit()
			// 	}
			// }
			continue
		default:
			continue

		}

	}

	fmt.Println("seeding completed sucess")
}

func init() {
	goBlueCmd.AddCommand(seedCmd)

}
