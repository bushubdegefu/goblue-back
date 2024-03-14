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
	ID          uint      `validate:"required" json:"id"`
	Name        string    `validate:"required" json:"name"`
	Description string    `validate:"required" json:"description"`
	Active      bool      `validate:"required" json:"active"`
	Roles       []RoleGet `json:"roles,omitempty"`
}

type AppsMatrix struct {
	Role     string   `json:"role"`
	Features []string `json:"role_features,omitempty"`
}

type AppsDropDown struct {
	ID   uint   `validate:"required" json:"id"`
	Name string `validate:"required" json:"name"`
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
	var apps_get []AppGet
	result, err := common.PaginationPureModel(db, models.App{}, []models.App{}, uint(Page), uint(Limit))
	if err != nil {
		return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all apps.",
			Data:    "something",
		})
	}
	mapstructure.Decode(result.Items, &apps_get)
	result.Items = apps_get
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

// GetApp Features is a function to get a Apps by ID
// @Summary Get Features
// @Description Get App Features By ID
// @Tags App
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "App ID"
// @Success 200 {object} common.ResponseHTTP{data=[]AppsMatrix}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /appsmatrix/{id} [get]
func GetAppMatrix(contx *fiber.Ctx) error {
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
	if res := db.Model(&models.App{}).Preload(clause.Associations).Preload("Roles.Features").Preload(clause.Associations).Where("id = ?", id).First(&apps); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	var matrix_list []AppsMatrix
	for i := range apps.Roles {

		var app_matrix AppsMatrix
		role_key := apps.Roles[i].Name
		// fmt.Println(role_key)
		app_matrix.Role = role_key
		for j := range apps.Roles[i].Features {
			feature_append := apps.Roles[i].Features[j].Name
			app_matrix.Features = append(app_matrix.Features, feature_append)
			// app_matrix.RoleFeatures[role_key] = append(app_matrix.RoleFeatures[role_key], feature_append)
			// fmt.Println(feature_append)
		}

		matrix_list = append(matrix_list, app_matrix)
	}

	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got app Role Matrix",
		Data:    &matrix_list,
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
	app.Active = posted_app.Active

	tx := db.Begin()
	// add  data using transaction if values are valid
	if err := tx.Create(&app).Error; err != nil {
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
		Data:    &app,
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
	app.ID = uint(id)
	tx := db.Begin()
	if err := db.Model(&app).UpdateColumns(*patch_app).Update("active", patch_app.Active).Error; err != nil {
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
		Data:    &app,
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
	db.Model(&app).Association("Roles").Clear()
	db.Delete(&app)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a app.",
		Data:    &app,
	})
}

// Get Feature Dropdown only active roles
// @Summary Get FeatureDropDown
// @Description Get FeatureDropDown
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} common.ResponseHTTP{data=[]FeatureDropDown}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /appsdrop [get]
func GetDropApps(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var apps_drop []AppsDropDown
	if res := db.Model(&models.App{}).Where("active = ?", true).Find(&apps_drop); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got response",
		Data:    &apps_drop,
	})
}
