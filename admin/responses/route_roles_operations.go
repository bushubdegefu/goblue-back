package responses

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

// Add Role to Route
// @Summary Add Role to Route
// @Description Add Route Role
// @Tags RouteRole Operations
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param route_id path int true "Route ID"
// @Param role_id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/routerole/{route_id}/{role_id} [post]
func AddRouteRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	// validate path params
	route_id, err := strconv.Atoi(contx.Params("route_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// validate path params
	role_id, err := strconv.Atoi(contx.Params("role_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// startng create transaction
	route_role := new(models.RouteRoles)
	route_role.RouteResponseID = uint(route_id)
	route_role.RoleID = uint(role_id)
	tx := db.Begin()
	if err := db.Model(&route_role).Create(&route_role).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	tx.Commit()

	// fetching added role
	var role RoleGet
	if res := db.Model(&models.Role{}).Where("id = ?", role_id).First(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Creating a role Route.",
		Data:    role,
	})
}

// Delete Role to Route
// @Summary Add Role
// @Description Delete Route Role
// @Tags RouteRole Operations
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param route_id path int true "Route ID"
// @Param role_id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/routerole/{route_id}/{role_id} [delete]
func DeleteRouteRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	// validate path params
	route_id, err := strconv.Atoi(contx.Params("route_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "error parssing at route_id: " + err.Error(),
			Data:    nil,
		})
	}

	role_id, err := strconv.Atoi(contx.Params("role_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "error parssing at role_id: " + err.Error(),
			Data:    nil,
		})
	}

	// instance for row to be deleted
	// route_role := new(models.RouteRoles)

	// starting transaction
	tx := db.Begin()
	if err := db.Where("route_response_id = ?", route_id).Where("role_id = ?", role_id).Delete(&models.RouteRoles{}).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}

	tx.Commit()

	var role RoleGet
	if res := db.Model(&models.Role{}).Where("id = ?", role_id).First(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Removing a role from route.",
		Data:    role,
	})
}
