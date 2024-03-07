package responses

// https://morkid.github.io/paginate/
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

type PageGet struct {
	ID          uint          `gorm:"primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name        string        `gorm:"type:string; constraint:not null;" json:"name,omitempty"`
	App         string        `gorm:"type:string; constraint:not null;" json:"App,omitempty"`
	Active      bool          `gorm:"default:true; constraint:not null;" json:"active" `
	Description string        `gorm:"type:string;" json:"description,omitempty"`
	Roles       []models.Role `gorm:"many2many:page_roles" json:"roles,omitempty"`
}

// Page Post model info
// @Description Page type information
// @Description Contains id name and description
type PagePost struct {
	Name        string `validate:"required" json:"name,omitempty"`
	Description string `validate:"required"  json:"description,omitempty"`
}

// Page Post model info
// @Description Page type information
// @Description Contains id name and description
type PagePatch struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Active      bool   `json:"active,omitempty"`
}

// GetPages is a function to get a Pages by ID
// @Summary Get Pages
// @Description Get Pages
// @Tags Page
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security Refresh
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]PageGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /pages [get]
func GetPages(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	Page, _ := strconv.Atoi(contx.Query("page"))
	Limit, _ := strconv.Atoi(contx.Query("size"))
	if Page == 0 || Limit == 0 {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Not Allowed, Bad request",
			Data:    nil,
		})
	}

	result, err := common.Pagination(db, models.Page{}, []models.Page{}, uint(Page), uint(Limit))

	if err != nil {
		return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all pages.",
			Data:    err.Error(),
		})
	}
	return contx.JSON(result)

}

// GetPageByID is a function to get a Pages by ID
// @Summary Get Page by ID
// @Description Get page by ID
// @Tags Page
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Page ID"
// @Success 200 {object} common.ResponseHTTP{data=PageGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /pages/{id} [get]
func GetPageID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var page models.Page
	if res := db.Model(&models.Page{}).Preload(clause.Associations).Where("id = ?", id).First(&page); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one page.",
		Data:    &page,
	})
}

// Get Roles of Page By ID
// @Summary Get Page Roles by ID
// @Description Get Roles by page by ID
// @Tags Page
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page_id path int true "Page ID"
// @Success 200 {object} common.ResponseHTTP{data=PageGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /pagesroles/{page_id} [get]
func GetPageRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	page_id, err := strconv.Atoi(contx.Params("page_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var roles []models.Role
	var page models.Page
	if res := db.Model(&models.Page{}).Where("id = ?", page_id).Find(&page); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	db.Model(&page).Association("Roles").Find(&roles)

	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one route.",
		Data:    &roles,
	})
}

// Add Page to data
// @Summary Add a new Page
// @Description Add Page
// @Tags Page
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page body PagePost true "Add Page"
// @Success 200 {object} common.ResponseHTTP{data=PagePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /pages [post]
func PostPage(contx *fiber.Ctx) error {

	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	posted_page := new(PagePost)

	//first parse post data
	if err := contx.BodyParser(&posted_page); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "parsing error : " + err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_page); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "validation error:" + err.Error(),
			Data:    nil,
		})
	}
	page := new(models.Page)
	page.Name = posted_page.Name
	page.Description = posted_page.Description
	tx := db.Begin()
	// add  data using transaction if values are valid
	if err := tx.Create(&page).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Page Creation Failed",
			Data:    err,
		})
	}
	tx.Commit()

	// return data if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a page.",
		Data:    page,
	})
}

// Patch Page to data
// @Summary Patch Page
// @Description Patch Page
// @Tags Page
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page body PagePost true "Patch Page"
// @Param id path int true "Page ID"
// @Success 200 {object} common.ResponseHTTP{data=PagePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /pages/{id} [patch]
func PatchPage(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()
	// validate path params
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// validate data struct
	// first parsing
	patch_page := new(PagePatch)
	if err := contx.BodyParser(&patch_page); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// then validating
	if err := validate.Struct(patch_page); err != nil {
		fmt.Println(err.Error())
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// startng update transaction
	page := new(models.Page)
	page.ID = uint(id)
	tx := db.Begin()
	if err := db.Model(&page).UpdateColumns(*patch_page).Update("active", patch_page.Active).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	tx.Commit()

	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Updating a page.",
		Data:    page,
	})
}

// DeletePages function removes a page by ID
// @Summary Remove Page by ID
// @Description Remove page by ID
// @Tags Page
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Page ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /pages/{id} [delete]
func DeletePage(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var page models.Page
	// validate path params
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// perform delete operation if the object exists
	tx := db.Begin()
	if err := db.Where("id = ?", id).First(&page).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	db.Delete(&page)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a page.",
		Data:    page,
	})
}
