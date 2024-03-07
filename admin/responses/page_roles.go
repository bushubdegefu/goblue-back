package responses

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

// Add Role to Page
// @Summary Add Role to Page
// @Description Add Page Role
// @Tags Page
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page_id path int true "Page ID"
// @Param role_id path int true "Role ID"
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /pagerole/{page_id}/{role_id} [post]
func AddPageRoles(contx *fiber.Ctx) error {
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
	role_id, err := strconv.Atoi(contx.Params("role_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// fetching role to be added
	var role models.Role
	role.ID = uint(role_id)
	if res := db.Find(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	//  appending assocation
	var page models.Page
	page.ID = uint(page_id)
	if err := db.Find(&page); err.Error != nil {
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err.Error.Error(),
		})
	}

	tx := db.Begin()
	if err := db.Model(&page).Association("Roles").Append(&role); err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Appending Role Failed",
			Data:    err.Error(),
		})
	}
	tx.Commit()

	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Creating a role Page.",
		Data:    role,
	})
}

// Delete Role to Page
// @Summary Add Role
// @Description Delete Page Role
// @Tags Page
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page_id path int true "Page ID"
// @Param role_id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /pagerole/{page_id}/{role_id} [delete]
func DeletePageRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	// validate path params
	page_id, err := strconv.Atoi(contx.Params("page_id"))
	if err != nil || page_id == 0 {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	role_id, err := strconv.Atoi(contx.Params("role_id"))
	if err != nil || role_id == 0 {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// fetching role to be deleted
	var role models.Role
	role.ID = uint(role_id)
	if res := db.Find(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// fettchng page
	var page models.Page
	page.ID = uint(page_id)
	if err := db.Find(&page); err.Error != nil {
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err.Error.Error(),
		})
	}

	// removing role
	tx := db.Begin()
	if err := db.Model(&page).Association("Roles").Delete(&role); err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNonAuthoritativeInfo).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Please Try Again Something Unexpected Happened",
			Data:    err.Error(),
		})
	}

	tx.Commit()

	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Removing a role from page.",
		Data:    role,
	})
}
