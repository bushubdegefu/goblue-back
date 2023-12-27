package responses

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

// Add Role to User
// @Summary Add Role to User
// @Description Add User Role
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param role_id path int true "Role ID"
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /userrole/{user_id}/{role_id} [post]
func AddUserRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	// validate path params
	user_id, err := strconv.Atoi(contx.Params("user_id"))
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
	if res := db.Model(&models.Role{}).Where("id = ?", role_id).First(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	//  appending assocation
	var user models.User
	if err := db.Model(&models.User{}).Where("id = ?", user_id).First(&user); err.Error != nil {
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err.Error.Error(),
		})
	}

	tx := db.Begin()
	if err := db.Model(&user).Association("Roles").Append(&role); err != nil {
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
		Message: "Success Creating a role User.",
		Data:    role,
	})
}

// Delete Role to User
// @Summary Add Role
// @Description Delete User Role
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param role_id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /userrole/{user_id}/{role_id} [delete]
func DeleteUserRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	// validate path params
	user_id, err := strconv.Atoi(contx.Params("user_id"))
	if err != nil || user_id == 0 {
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
	if res := db.Model(&models.Role{}).Where("id = ?", role_id).First(&role); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	// fettchng user
	var user models.User
	if err := db.Model(&models.User{}).Where("id = ?", user_id).First(&user); err.Error != nil {
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err.Error.Error(),
		})
	}

	// removing role
	tx := db.Begin()
	if err := db.Model(&user).Association("Roles").Delete(&role); err != nil {
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
		Message: "Success Removing a role from user.",
		Data:    role,
	})
}
