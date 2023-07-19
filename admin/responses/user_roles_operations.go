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
// @Tags UserRole Operations
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param role_id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/userrole/{user_id}/{role_id} [post]
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

	// startng create transaction
	user_role := new(models.UserRoles)
	user_role.UserID = uint(user_id)
	user_role.RoleID = uint(role_id)
	tx := db.Begin()
	if err := db.Model(&user_role).Create(&user_role).Error; err != nil {
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
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Creating a role User.",
		Data:    role,
	})
}

// Delete Role to User
// @Summary Add Role
// @Description Delete User Role
// @Tags UserRole Operations
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param role_id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/userrole/{user_id}/{role_id} [delete]
func DeleteUserRoles(contx *fiber.Ctx) error {
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

	role_id, err := strconv.Atoi(contx.Params("role_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// instance for row to be deleted
	// user_role := new(models.UserRoles)

	// starting transaction
	tx := db.Begin()
	if err := db.Where("user_id = ?", user_id).Where("role_id = ?", role_id).Delete(&models.UserRoles{}).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	// db.Delete(&user_role)

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
		Message: "Success Removing a role from user.",
		Data:    role,
	})
}
