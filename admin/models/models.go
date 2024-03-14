package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"semay.com/admin/database"
)

type Role struct {
	ID          uint          `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string        `gorm:"not null; unique;" json:"name,omitempty"`
	Description string        `gorm:"not null; unique;" json:"description,omitempty"`
	Active      bool          `gorm:"default:true; constraint:not null;" json:"active" `
	Users       []User        `gorm:"many2many:user_roles; constraint:OnUpdate:CASCADE; OnDelete:CASCADE;" json:"users,omitempty"`
	Features    []Feature     `gorm:"foreignkey:RoleID; constraint:OnUpdate:CASCADE; OnDelete:SET NULL;" json:"features,omitempty"`
	Pages       []Page        `gorm:"many2many:page_roles; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pages,omitempty"`
	AppID       sql.NullInt64 `gorm:"foreignkey:AppID OnDelete:SET NULL" json:"app,omitempty" swaggertype:"number"`
}

type EndPoint struct {
	ID          uint          `gorm:"primaryKey;autoIncrement:true"`
	Name        string        `gorm:"type:string; unique;" json:"name,omitempty"`
	RoutePaths  string        `gorm:"type:string;" json:"route_path,omitempty"`
	Method      string        `gorm:"type:string;" json:"method,omitempty"`
	Description string        `gorm:"type:string;" json:"description,omitempty"`
	FeatureID   sql.NullInt64 `gorm:"foreignkey:FeatureID default:NULL;,OnDelete:SET NULL;" json:"feature_id,omitempty" swaggertype:"number"`
}

type Page struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name        string `gorm:"type:string; constraint:not null; unique;" json:"name,omitempty"`
	Active      bool   `gorm:"default:true; constraint:not null;" json:"active,omitempty" `
	Description string `gorm:"type:string;" json:"description,omitempty"`
	Roles       []Role `gorm:"many2many:page_roles; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"roles,omitempty"`
}

type App struct {
	ID          uint      `gorm:"primaryKey;autoIncrement:true"`
	Name        string    `gorm:"type:string; unique; constraint:not null;" json:"name,omitempty"`
	UUID        uuid.UUID `gorm:"constraint:not null; type:uuid;" json:"uuid"`
	Active      bool      `gorm:"default:true; constraint:not null;" json:"active,omitempty" `
	Description string    `gorm:"type:string;" json:"description,omitempty"`
	Roles       []Role    `gorm:"association_foreignkey:AppID constraint:OnUpdate:SET NULL OnDelete:SET NULL" json:"roles,omitempty"`
}

func (app *App) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	app.UUID = uuid.New()
	return
}

type Feature struct {
	ID          uint          `gorm:"primaryKey;autoIncrement:true"`
	Name        string        `gorm:"type:string; unique;" json:"name,omitempty"`
	Description string        `gorm:"type:string;" json:"description,omitempty"`
	Active      bool          `gorm:"default:true; constraint:not null;" json:"active,omitempty" `
	Endpoints   []EndPoint    `gorm:"association_foreignkey:FeatureID constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"endpoints,omitempty"`
	RoleID      sql.NullInt64 `gorm:"foreignkey:RoleID; constraint:OnDelete:SET NULL;" json:"role,omitempty" swaggertype:"number"`
}

type User struct {
	ID             uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	UUID           uuid.UUID `gorm:"constraint:not null; type:uuid;" json:"uuid"`
	Email          string    `gorm:"constraint:not null; unique;" json:"email"`
	Password       string    `gorm:"constraint:not null;" json:"password"`
	DateRegistered time.Time `gorm:"constraint:not null; default:current_timestamp;" json:"date_registered,omitempty"`
	Disabled       bool      `gorm:"constraint:not null; default:false;" json:"disabled,omitempty"`
	Roles          []Role    `gorm:"many2many:user_roles; constraint:OnUpdate:CASCADE; OnDelete:CASCADE;" json:"roles,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.UUID = uuid.New()
	user.Password = hashfunc(user.Password)
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

type SessionData struct {
	Token     string    `gorm:"constraint:not null;" json:"token"`
	TimeStamp time.Time `gorm:"constraint:not null; default:current_timestamp;" json:"signed_time"`
}

type BlobPicture struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	BlobPicture []byte `json:"blob_picture"`
}

type BlobVideo struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"index" json:"name"`
	BlobVideo []byte `json:"blob_video"`
}

type JWTSalt struct {
	ID    uint   `gorm:"primaryKey;autoIncrement:true"`
	SaltA string `gorm:"constraint:not null;" json:"salt_a"`
	SaltB string `gorm:"constraint:not null;" json:"salt_b"`
}

func InitDatabase() {
	database := database.ReturnSession()
	fmt.Println("Connection Opened to Database")
	if err := database.AutoMigrate(
		&User{},
		&App{},
		&Role{},
		&Page{},
		&EndPoint{},
		&Feature{},
		&SessionData{},
		&SiteData{},
		&BlobPicture{},
		&BlobVideo{},
		&JWTSalt{},
	); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Database Migrated")
}
