package responses

// https://morkid.github.io/paginate/
import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
)

type RoleGet struct {
	ID          uint   `validate:"required"`
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

// Role Post model info
// @Description Role type information
// @Description Contains id name and description
type RolePost struct {
	Name        string `json:"name" example:"superuser"`
	Description string `json:"description" example:"Devloper Mode Acecss"`
}

// GetRoles is a function to get a Roles by ID
// @Summary Get Roles
// @Description Get Roles
// @Tags Role
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security Refresh
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]RoleGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /roles [get]
func GetRoles(contx *fiber.Ctx) error {
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

	result, err := common.PaginationPureModel(db, models.Role{}, []RoleGet{}, uint(Page), uint(Limit))
	if err != nil {
		return contx.JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all roles.",
			Data:    "something",
		})
	}
	return contx.JSON(result)

}

// GetRoleByID is a function to get a Roles by ID
// @Summary Get Role by ID
// @Description Get role by ID
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RoleGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /roles/{id} [get]
func GetRolesID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var roles RoleGet
	if res := db.Model(&models.Role{}).Where("id = ?", id).First(&roles); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one role.",
		Data:    &roles,
	})
}

// Add Role to data
// @Summary Add a new Role
// @Description Add Role
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param role body RolePost true "Add Role"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /roles [post]
func PostRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	posted_role := new(RolePost)

	//first parse post data
	if err := contx.BodyParser(&posted_role); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_role); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	role := new(models.Role)
	role.Name = posted_role.Name
	role.Description = posted_role.Description
	tx := db.Begin()
	// add  data using transaction if values are valid
	// if err := tx.Create(&role).Error; err != nil {
	if err := tx.Model(&role).Create(&role).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Role Creation Failed",
			Data:    err,
		})
	}
	tx.Commit()

	// return data if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a role.",
		Data:    role,
	})
}

// Patch Role to data
// @Summary Patch Role
// @Description Patch Role
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param role body RolePost true "Patch Role"
// @Param id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{data=RolePost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /roles/{id} [patch]
func PatchRoles(contx *fiber.Ctx) error {
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
	patch_role := new(RolePost)
	if err := contx.BodyParser(&patch_role); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// then validating
	if err := validate.Struct(patch_role); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// startng update transaction
	role := new(models.Role)
	tx := db.Begin()
	if err := db.Model(&role).Where("id = ?", id).First(&role).UpdateColumns(*patch_role).Error; err != nil {
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
		Message: "Success Updating a role.",
		Data:    role,
	})
}

// DeleteRoles function removes a role by ID
// @Summary Remove Role by ID
// @Description Remove role by ID
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /roles/{id} [delete]
func DeleteRoles(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var role models.Role
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
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	db.Delete(&role)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a role.",
		Data:    role,
	})
}
