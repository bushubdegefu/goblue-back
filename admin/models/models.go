package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"semay.com/admin/database"
	"semay.com/config"
	"semay.com/utils"
)

type Role struct {
	ID          uint            `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string          `gorm:"not null; unique;" json:"name,omitempty"`
	Description string          `gorm:"not null; unique;" json:"description,omitempty"`
	Users       []User          `gorm:"many2many:user_roles" json:"users,omitempty"`
	Routes      []RouteResponse `gorm:"many2many:route_roles" json:"routes,omitempty"`
}

type RouteResponse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `gorm:"type:string;" json:"name,omitempty"`
	RoutePaths  string `gorm:"type:string;" json:"route_path,omitempty"`
	Description string `gorm:"type:string;" json:"description,omitempty"`
	Pages       []Page `gorm:"many2many:page_routes" json:"pages,omitempty"`
	Roles       []Role `gorm:"many2many:route_roles" json:"roles,omitempty"`
}

type Page struct {
	ID          uint            `gorm:"primaryKey;autoIncrement:true"`
	Name        string          `gorm:"type:string; constraint:not null;" json:"name,omitempty"`
	App         string          `gorm:"type:string; constraint:not null;" json:"App,omitempty"`
	Active      bool            `gorm:"default:true; constraint:not null;" json:"active,omitempty" `
	Description string          `gorm:"type:string;" json:"description,omitempty"`
	Routes      []RouteResponse `gorm:"many2many:page_routes" json:"routes,omitempty"`
}
type User struct {
	ID             uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	UUID           uuid.UUID `gorm:"constraint:not null; type:uuid;" json:"uuid"`
	Email          string    `gorm:"constraint:not null;" json:"email"`
	Password       string    `gorm:"constraint:not null;" json:"password"`
	DateRegistered time.Time `gorm:"constraint:not null; default:current_timestamp;" json:"date_registered"`
	Disabled       bool      `gorm:"constraint:not null; default:true;" json:"disabled"`
	Roles          []Role    `gorm:"many2many:user_roles;" json:"roles"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.UUID = uuid.New()
	user.Password = utils.HashFunc(user.Password)
	return
}

type SiteData struct {
	ID             uint    `gorm:"primaryKey;autoIncrement:true"`
	RemoteAdd      string  `gorm:"type:varchar(128)"`
	AccessedRoute  string  `gorm:"type:varchar(300)"`
	Method         string  `gorm:"type:varchar(10)"`
	ResponseTime   float64 `gorm:"type:float"`
	ResponseStatus float64 `gorm:"type:int"`
}

type UserRoles struct {
	RoleID uint `gorm:"association_foreignkey:RoleID"`
	UserID uint `gorm:"association_foreignkey:UserID"`
}

type RouteRoles struct {
	RoleID          uint `gorm:"association_foreignkey:RoleID"`
	RouteResponseID uint `gorm:"association_foreignkey:UserID"`
}

type PageRoutes struct {
	PageID          uint `gorm:"association_foreignkey:RoleID"`
	RouteResponseID uint `gorm:"association_foreignkey:UserID"`
}

func InitDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open(config.Config("SQLITE_URI")))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	// database.DBConn.AutoMigrate(&book.Book{})

	if err := database.DBConn.AutoMigrate(
		&User{},
		&Role{},
		&Page{},
		&RouteResponse{},
		&SiteData{},
	); err != nil {
		log.Fatalln(err)
	}
	// database.DBConn.Model(&UserRoles{}).AddForeignKey("role_id", "roles(role_id)", "CASCADE", "CASCADE")
	// database.DBConn.Model(&UserRoles{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	fmt.Println("Database Migrated")
}
