package responses

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

// Add Route to Page
// @Summary Add Route to Page
// @Description Add Page Route
// @Tags PageRoute Operations
// @Accept json
// @Produce json
// @Param page_id path int true "Page ID"
// @Param route_id path int true "Route ID"
// @Success 200 {object} common.ResponseHTTP{data=RoutePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/pageroute/{page_id}/{route_id} [post]
func AddPageRoutes(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	// validate path params
	page_id, err := strconv.Atoi(contx.Params("page_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// validate path params
	route_id, err := strconv.Atoi(contx.Params("route_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// startng create transaction
	page_route := new(models.PageRoutes)
	page_route.PageID = uint(page_id)
	page_route.RouteResponseID = uint(route_id)
	tx := db.Begin()
	if err := db.Model(&page_route).Create(&page_route).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	tx.Commit()

	// fetching added route
	var route models.RouteResponse
	if res := db.Model(&models.RouteResponse{}).Where("id = ?", route_id).First(&route); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Creating a Route to Page.",
		Data:    route,
	})
}

// Delete Route to Page
// @Summary Add Route
// @Description Delete Page Route
// @Tags PageRoute Operations
// @Accept json
// @Produce json
// @Param page_id path int true "Page ID"
// @Param route_id path int true "Route ID"
// @Success 200 {object} common.ResponseHTTP{data=RoutePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/pageroute/{page_id}/{route_id} [delete]
func DeletePageRoutes(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	// validate path params
	page_id, err := strconv.Atoi(contx.Params("page_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	route_id, err := strconv.Atoi(contx.Params("route_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// starting transaction
	tx := db.Begin()
	if err := db.Where("page_id = ?", page_id).Where("route_response_id = ?", route_id).Delete(&models.PageRoutes{}).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}

	tx.Commit()

	var route RouteGet
	if res := db.Model(&models.RouteResponse{}).Where("id = ?", route_id).First(&route); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Removing a route from page.",
		Data:    route,
	})
}
