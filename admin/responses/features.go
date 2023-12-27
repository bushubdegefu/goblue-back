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

type FeatureGet struct {
	ID          uint          `validate:"required" json:"id"`
	Name        string        `validate:"required" json:"name"`
	Description string        `validate:"required" json:"description"`
	Active      bool          `validate:"required" json:"active"`
	Endpoints   []EndPointGet `json:"endpoints,omitempty"`
}

type FeatureDropDown struct {
	ID   uint   `validate:"required" json:"id"`
	Name string `validate:"required" json:"name"`
}

// Feature Post model info
// @Description Feature type information
// @Description Contains id name and description
type FeaturePost struct {
	Name        string `json:"name" example:"superuser"`
	Description string `json:"description" example:"Devloper Mode Acecss"`
	Active      bool   `json:"active" example:"true"`
}

// GetFeatures is a function to get a Features by ID
// @Summary Get Features
// @Description Get Features
// @Tags Feature
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security Refresh
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]FeatureGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /features [get]
func GetFeatures(contx *fiber.Ctx) error {
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

	var features_get []FeatureGet
	result, err := common.PaginationPureModel(db, models.Feature{}, []models.Feature{}, uint(Page), uint(Limit))
	if err != nil {
		return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all features.",
			Data:    "something",
		})
	}

	mapstructure.Decode(result.Items, &features_get)
	result.Items = features_get
	return contx.Status(http.StatusOK).JSON(result)

}

// GetFeatureByID is a function to get a Features by ID
// @Summary Get Feature by ID
// @Description Get feature by ID
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Feature ID"
// @Success 200 {object} common.ResponseHTTP{data=FeatureGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /features/{id} [get]
func GetFeaturesID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var features models.Feature
	var features_get FeatureGet
	if res := db.Model(&models.Feature{}).Preload(clause.Associations).Where("id = ?", id).First(&features); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	mapstructure.Decode(features, &features_get)
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one feature.",
		Data:    &features_get,
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
// @Router /featuredrop [get]
func GetDropFeatures(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var features_drop []FeatureDropDown
	if res := db.Model(&models.Feature{}).Where("active = ?", true).Find(&features_drop); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one role.",
		Data:    &features_drop,
	})
}

// Add Feature to data
// @Summary Add a new Feature
// @Description Add Feature
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param feature body FeaturePost true "Add Feature"
// @Success 200 {object} common.ResponseHTTP{data=FeaturePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /features [post]
func PostFeatures(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	posted_feature := new(FeaturePost)

	//first parse post data
	if err := contx.BodyParser(&posted_feature); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_feature); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	feature := new(models.Feature)
	feature.Name = posted_feature.Name
	feature.Description = posted_feature.Description
	tx := db.Begin()
	// add  data using transaction if values are valid
	// if err := tx.Create(&feature).Error; err != nil {
	if err := tx.Model(&feature).Create(&feature).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Feature Creation Failed",
			Data:    err,
		})
	}
	tx.Commit()

	// return data if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a feature.",
		Data:    feature,
	})
}

// Patch Feature to data
// @Summary Patch Feature
// @Description Patch Feature
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param feature body FeaturePost true "Patch Feature"
// @Param id path int true "Feature ID"
// @Success 200 {object} common.ResponseHTTP{data=FeaturePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /features/{id} [patch]
func PatchFeatures(contx *fiber.Ctx) error {
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
	patch_feature := new(FeaturePost)
	if err := contx.BodyParser(&patch_feature); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// then validating
	if err := validate.Struct(patch_feature); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// startng update transaction
	feature := new(models.Feature)
	tx := db.Begin()
	if err := db.Model(&feature).Where("id = ?", id).First(&feature).UpdateColumns(*patch_feature).Error; err != nil {
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
		Message: "Success Updating a feature.",
		Data:    feature,
	})
}

// Activate/Deactivate Feature to data
// @Summary Activate/Deactivate Feature
// @Description Activate/Deactivate Feature
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param feature_id path int true "Feature ID"
// @Param active query bool true "Active"
// @Success 200 {object} common.ResponseHTTP{data=FeaturePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /features/{feature_id} [put]
func ActivateDeactivateFeature(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	// validate path params
	feature_id, err := strconv.Atoi(contx.Params("feature_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	//  gettng query parm
	active := contx.QueryBool("active")
	// startng update transaction
	var feature models.Feature
	tx := db.Begin()
	if err := db.Model(&feature).Where("id = ?", feature_id).First(&feature).Update("active", active).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	tx.Commit()
	feature.Active = active
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Updating a feature.",
		Data:    feature,
	})
}

// DeleteFeatures function removes a feature by ID
// @Summary Remove Feature by ID
// @Description Remove feature by ID
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Feature ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /features/{id} [delete]
func DeleteFeatures(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var feature models.Feature
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
	if err := db.Where("id = ?", id).First(&feature).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	db.Delete(&feature)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a feature.",
		Data:    feature,
	})
}

// Add Feature Role
// @Summary Add Feature Role
// @Description Add Feature Role
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param feature_id path int true "Feature ID"
// @Param role_id query int true "Role ID"
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /featuresrole/{feature_id} [patch]
func AddFeatureRole(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	// validate path params
	feature_id, err := strconv.Atoi(contx.Params("feature_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// fetching role to be added
	var feature models.Feature
	if res := db.Model(&models.Feature{}).Where("id = ?", feature_id).First(&feature); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	// startng update transaction
	role_id := contx.QueryInt("role_id")

	var role models.Role
	if res := db.Model(&models.Role{}).Where("id = ?", role_id).First(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	tx := db.Begin()
	//  Adding one to many Relation
	if err := db.Model(&role).Association("Features").Append(&feature); err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Error Adding Record",
			Data:    err.Error(),
		})
	}
	tx.Commit()

	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Adding a Role to Feature.",
		Data:    feature,
	})
}

// Delete Feature Role
// @Summary Delete Feature Role
// @Description Delete Feature Role
// @Tags Feature
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param feature_id path int true "Feature ID"
// @Param role_id query int true "Role ID"
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /featuresrole/{feature_id} [delete]
func DeleteFeatureRole(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	// validate path params
	feature_id, err := strconv.Atoi(contx.Params("feature_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// fetching role to be added
	var feature models.Feature
	if res := db.Model(&models.Feature{}).Where("id = ?", feature_id).First(&feature); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// startng update transaction
	role_id := contx.QueryInt("role_id")
	var role models.Role
	if res := db.Model(&models.Role{}).Where("id = ?", role_id).First(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// Removing Role From Feature
	tx := db.Begin()
	if err := db.Model(&role).Association("Features").Delete(&feature); err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err.Error(),
		})
	}
	tx.Commit()

	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Deleteing a Role From Feature.",
		Data:    feature,
	})
}
