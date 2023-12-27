package responses

// https://morkid.github.io/paginate/
import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

// EndPoint Get model info
// @Description EndPoint type information
// @Description Contains id name and description
type EndPointGet struct {
	ID          uint   `json:"id,omitempty"`
	RoutePath   string `json:"route_path,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type EndPointDropDown struct {
	ID   uint   `validate:"required" json:"id"`
	Name string `validate:"required" json:"name"`
}

// EndPoint Post model info
// @Description EndPoint type information
// @Description Contains id name and description
type EndPointPost struct {
	Name        string `validate:"required" json:"name,omitempty"  example:"get users"`
	RoutePath   string `validate:"required" json:"route_path,omitempty"  example:"/route/path"`
	Description string `validate:"required" json:"description" example:"Fetchs user list"`
}

// GetEndpoints is a function to get a Endpoints by ID
// @Summary Get Endpoints
// @Description Get Endpoints
// @Tags EndPoints
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security Refresh
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]EndPointGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /endpoints [get]
func GetEndPointResponse(contx *fiber.Ctx) error {
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

	result, err := common.PaginationPureModel(db, models.EndPoint{}, []EndPointGet{}, uint(Page), uint(Limit))
	if err != nil {
		return contx.JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all routes.",
			Data:    err.Error(),
		})
	}
	return contx.Status(http.StatusOK).JSON(result)

}

// GetEndPointByID is a function to get a EndPoints by ID
// @Summary Get EndPoint by ID
// @Description Get route by ID
// @Tags EndPoints
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "EndPoint ID"
// @Success 200 {object} common.ResponseHTTP{data=EndPointGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /endpoints/{id} [get]
func GetEndPointsID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var routes models.EndPoint
	if res := db.Model(&models.EndPoint{}).Preload(clause.Associations).Where("id = ?", id).First(&routes); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one route.",
		Data:    &routes,
	})
}

// Get EndPoint Dropdown only active roles
// @Summary Get EndPointDropDown
// @Description Get EndPointDropDown
// @Tags EndPoint
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} common.ResponseHTTP{data=[]EndPointDropDown}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /endpointdrop [get]
func GetDropEndPoints(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var features_drop []EndPointDropDown
	if res := db.Model(&models.EndPoint{}).Find(&features_drop); res.Error != nil {
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

// Add EndPoint to data
// @Summary Add a new EndPoint
// @Description Add EndPoint
// @Tags EndPoints
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param route body EndPointPost true "Add EndPoint"
// @Success 200 {object} common.ResponseHTTP{data=EndPointPost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /endpoints [post]
func PostEndPoint(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	posted_route := new(EndPointPost)

	//first parse post data
	if err := contx.BodyParser(&posted_route); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_route); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	route := new(models.EndPoint)
	route.Name = posted_route.Name
	route.Description = posted_route.Description
	tx := db.Begin()
	// add  data using transaction if values are valid
	// if err := tx.Create(&route).Error; err != nil {
	if err := tx.Model(&route).Create(&route).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "EndPoint Creation Failed",
			Data:    err,
		})
	}
	tx.Commit()

	// return data if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a route.",
		Data:    route,
	})
}

// Patch EndPoint to data
// @Summary Patch EndPoint
// @Description Patch EndPoint
// @Tags EndPoints
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param route body EndPointPost true "Patch EndPoint"
// @Param id path int true "EndPoint ID"
// @Success 200 {object} common.ResponseHTTP{data=EndPointPost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /endpoints/{id} [patch]
func PatchEndPoint(contx *fiber.Ctx) error {
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
	patch_route := new(EndPointPost)
	if err := contx.BodyParser(&patch_route); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// then validating
	if err := validate.Struct(patch_route); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// startng update transaction
	route := new(models.EndPoint)
	tx := db.Begin()
	if err := db.Model(&route).Where("id = ?", id).First(&route).UpdateColumns(*patch_route).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	tx.Commit()

	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Updating a route.",
		Data:    route,
	})
}

// DeleteEndPoints function removes a route by ID
// @Summary Remove EndPoint by ID
// @Description Remove route by ID
// @Tags EndPoints
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "EndPoint ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /endpoints/{id} [delete]
func DeleteEndPoint(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var route models.EndPoint
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
	if err := db.Where("id = ?", id).First(&route).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	db.Delete(&route)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a route.",
		Data:    route,
	})
}

// Add Feature Endpoint
// @Summary Add Feature to Endpoint
// @Description Add Feature to Endpoint
// @Tags EndPoints
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param endpoint_id path int true "Feature ID"
// @Param feature_id query int true "Endpoint ID"
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /feature_endpoint/{endpoint_id} [patch]
func AddEndpointFeature(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	// validate path params
	endpoint_id, err := strconv.Atoi(contx.Params("endpoint_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// fetching Endpionts
	var endpoint models.EndPoint
	if res := db.Model(&models.EndPoint{}).Where("id = ?", endpoint_id).First(&endpoint); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// fetching role to be added
	feature_id := contx.QueryInt("feature_id")
	var feature models.Feature
	if res := db.Model(&models.Feature{}).Where("id = ?", feature_id).First(&feature); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// startng update transaction

	tx := db.Begin()
	//  Adding one to many Relation
	if err := db.Model(&feature).Association("Endpoints").Append(&endpoint); err != nil {
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
		Message: "Success Adding a Endpoint to Feature.",
		Data:    feature,
	})
}

// Delete Feature Endpoint
// @Summary Delete Feature Endpoint
// @Description Delete Feature Endpoint
// @Tags EndPoints
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param endpoint_id path int true "Feature ID"
// @Param feature_id query int true "Endpoint ID"
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /feature_endpoint/{endpoint_id} [delete]
func DeleteEndpointFeature(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	// validate path params
	endpoint_id, err := strconv.Atoi(contx.Params("endpoint_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Getting Endpoint
	var endpoint models.EndPoint
	if res := db.Model(&models.EndPoint{}).Where("id = ?", endpoint_id).First(&endpoint); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// fetching feature to be added
	var feature models.Feature
	feature_id := contx.QueryInt("feature_id")
	if res := db.Model(&models.Feature{}).Where("id = ?", feature_id).First(&feature); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// Removing Endpoint From Feature
	tx := db.Begin()
	if err := db.Model(&feature).Association("Endpoints").Delete(&endpoint); err != nil {
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
		Message: "Success Deleteing a Endpoint From Feature.",
		Data:    feature,
	})
}
