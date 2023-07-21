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

type RouteGet struct {
	ID          uint   `json:"id,omitempty"`
	RoutePaths  string `json:"RoutePath,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Route Post model info
// @Description Route type information
// @Description Contains id name and description
type RoutePost struct {
	Name        string `validate:"required" json:"name,omitempty"  example:"get users"`
	RoutePaths  string `validate:"required" json:"route_path,omitempty"  example:"/route/path"`
	Description string `validate:"required" json:"description" example:"Fetchs user list"`
}

// GetRoutes is a function to get a Routes by ID
// @Summary Get Routes
// @Description Get Routes
// @Tags Routes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security Refresh
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]RouteGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/routes [get]
func GetRouteResponse(contx *fiber.Ctx) error {
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

	result, err := common.PaginationPureModel(db, models.RouteResponse{}, []RouteGet{}, uint(Page), uint(Limit))
	if err != nil {
		return contx.JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all routes.",
			Data:    "something",
		})
	}
	return contx.Status(http.StatusOK).JSON(result)

}

// GetRouteByID is a function to get a Routes by ID
// @Summary Get Route by ID
// @Description Get route by ID
// @Tags Routes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Route ID"
// @Success 200 {object} common.ResponseHTTP{data=RouteGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/routes/{id} [get]
func GetRoutesID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var routes models.RouteResponse
	if res := db.Model(&models.RouteResponse{}).Preload(clause.Associations).Where("id = ?", id).First(&routes); res.Error != nil {
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

// Get Role of Routes By ID
// @Summary Get Route Roles by ID
// @Description Get route by ID
// @Tags Routes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param route_id path int true "RouteResponse ID"
// @Success 200 {object} common.ResponseHTTP{data=RouteGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/routerole/{route_id} [get]
func GetRouteRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	route_id, err := strconv.Atoi(contx.Params("route_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var roles []models.Role
	var route models.RouteResponse
	if res := db.Model(&models.RouteResponse{}).Where("id = ?", route_id).Find(&route); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	db.Model(&route).Association("Roles").Find(&roles)

	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one route.",
		Data:    &roles,
	})
}

// Add Route to data
// @Summary Add a new Route
// @Description Add Route
// @Tags Routes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param route body RoutePost true "Add Route"
// @Success 200 {object} common.ResponseHTTP{data=RoutePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/routes [post]
func PostRoute(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	posted_route := new(RoutePost)

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
	route := new(models.RouteResponse)
	route.Name = posted_route.Name
	route.Description = posted_route.Description
	tx := db.Begin()
	// add  data using transaction if values are valid
	// if err := tx.Create(&route).Error; err != nil {
	if err := tx.Model(&route).Create(&route).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Route Creation Failed",
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

// Patch Route to data
// @Summary Patch Route
// @Description Patch Route
// @Tags Routes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param route body RoutePost true "Patch Route"
// @Param id path int true "Route ID"
// @Success 200 {object} common.ResponseHTTP{data=RoutePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/routes/{id} [patch]
func PatchRoute(contx *fiber.Ctx) error {
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
	patch_route := new(RoutePost)
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
	route := new(models.RouteResponse)
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

// DeleteRoutes function removes a route by ID
// @Summary Remove Route by ID
// @Description Remove route by ID
// @Tags Routes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Route ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/routes/{id} [delete]
func DeleteRoute(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var route models.RouteResponse
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
