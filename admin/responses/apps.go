package responses

// https://morkid.github.io/paginate/
import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

type AppGet struct {
	ID          uint   `validate:"required"`
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Active      bool   `validate:"required"`
	Roles       []RoleGet
}

// App Post model info
// @Description App type information
// @Description Contains id name and description
type AppPost struct {
	Name        string `json:"name" example:"superuser"`
	Description string `json:"description" example:"Devloper Mode Acecss"`
	Active      bool   `json:"active" example:"true"`
}

// GetApps is a function to get a Apps by ID
// @Summary Get Apps
// @Description Get Apps
// @Tags App
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security Refresh
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]AppGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /apps [get]
func GetApps(contx *fiber.Ctx) error {
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

	result, err := common.PaginationPureModel(db, models.App{}, []AppGet{}, uint(Page), uint(Limit))
	if err != nil {
		return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all apps.",
			Data:    "something",
		})
	}
	return contx.Status(http.StatusOK).JSON(result)

}

// GetAppByID is a function to get a Apps by ID
// @Summary Get App by ID
// @Description Get app by ID
// @Tags App
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "App ID"
// @Success 200 {object} common.ResponseHTTP{data=AppGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /apps/{id} [get]
func GetAppsID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var apps models.App
	var response_app AppGet
	if res := db.Model(&models.App{}).Preload(clause.Associations).Where("id = ?", id).First(&apps); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	mapstructure.Decode(apps, &response_app)
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one app.",
		Data:    &response_app,
	})
}

// Add App to data
// @Summary Add a new App
// @Description Add App
// @Tags App
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param app body AppPost true "Add App"
// @Success 200 {object} common.ResponseHTTP{data=AppPost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /apps [post]
func PostApps(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	posted_app := new(AppPost)

	//first parse post data
	if err := contx.BodyParser(&posted_app); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_app); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	app := new(models.App)
	app.Name = posted_app.Name
	app.Description = posted_app.Description
	tx := db.Begin()
	// add  data using transaction if values are valid
	// if err := tx.Create(&app).Error; err != nil {
	if err := tx.Model(&app).Create(&app).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "App Creation Failed",
			Data:    err,
		})
	}
	tx.Commit()

	// return data if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a app.",
		Data:    app,
	})
}

// Patch App to data
// @Summary Patch App
// @Description Patch App
// @Tags App
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param app body AppPost true "Patch App"
// @Param id path int true "App ID"
// @Success 200 {object} common.ResponseHTTP{data=AppPost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /apps/{id} [patch]
func PatchApps(contx *fiber.Ctx) error {
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
	patch_app := new(AppPost)
	if err := contx.BodyParser(&patch_app); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// then validating
	if err := validate.Struct(patch_app); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// startng update transaction
	app := new(models.App)
	tx := db.Begin()
	if err := db.Model(&app).Where("id = ?", id).First(&app).UpdateColumns(*patch_app).Error; err != nil {
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
		Message: "Success Updating a app.",
		Data:    app,
	})
}

// DeleteApps function removes a app by ID
// @Summary Remove App by ID
// @Description Remove app by ID
// @Tags App
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "App ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /apps/{id} [delete]
func DeleteApps(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var app models.App
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
	if err := db.Where("id = ?", id).First(&app).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	db.Delete(&app)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a app.",
		Data:    app,
	})
}
