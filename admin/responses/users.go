package responses

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
	"semay.com/utils"
)

type UserGet struct {
	ID             uint      `json:"id,omitempty"`
	UUID           uuid.UUID `json:"uuid,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	DateRegistered time.Time `json:"date_registered,omitempty"`
	Disabled       bool      `json:"disabled,omitempty"`
}

type UserPost struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// GetUsers is a function to get a Users by ID
// @Summary Get Users
// @Description Get Users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]UserGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/users [get]
func GetUsers(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var Users []UserGet
	// // var pag_Users []models.User
	// Select `id`, `name` automatically when querying
	Page, _ := strconv.Atoi(contx.Query("page"))
	Limit, _ := strconv.Atoi(contx.Query("size"))
	if Page == 0 || Limit == 0 {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Not Allowed, Bad request",
			Data:    nil,
		})
	}

	result, err := common.PaginationPureModel(db, models.User{}, Users, uint(Page), uint(Limit))
	if err != nil {
		return contx.JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all Users.",
			Data:    "something",
		})
	}
	return contx.JSON(result)

}

// GetUserByID is a function to get a Users by ID
// @Summary Get User by ID
// @Description Get User by ID
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} common.ResponseHTTP{data=UserGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/users/{id} [get]
func GetUsersID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	id, err := strconv.Atoi(contx.Params("id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	var users models.User
	if res := db.Model(&models.User{}).Preload(clause.Associations).Where("id = ?", id).First(&users); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one User.",
		Data:    &users,
	})
}

// Get Roles of User By ID
// @Summary Get User Roles by ID
// @Description Get User Roles by ID
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} common.ResponseHTTP{data=[]RoleGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/userrole/{user_id} [get]
func GetUsersRolesByID(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	//validate user id
	user_id, err := strconv.Atoi(contx.Params("user_id"))
	if err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var roles []RoleGet
	var user models.User
	if res := db.Model(&models.User{}).Where("id = ?", user_id).Find(&user); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	db.Model(&user).Association("Roles").Find(&roles)

	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one route.",
		Data:    &roles,
	})
}

// Add Users to data
// @Summary Add a new Users
// @Description Add Users
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param User body UserPost true "Add User"
// @Success 200 {object} common.ResponseHTTP{data=UserPost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/users [post]
func PostUsers(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	posted_User := new(UserPost)

	//first parse post data
	if err := contx.BodyParser(&posted_User); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_User); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	User := new(models.User)
	User.Email = posted_User.Email
	User.Password = posted_User.Password
	tx := db.Begin()
	// add  data using transaction if values are valid
	// if err := tx.Create(&User).Error; err != nil {
	if err := tx.Model(&User).Create(&User).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "User Creation Failed",
			Data:    err,
		})
	}
	tx.Commit()

	// return data if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a User.",
		Data:    User,
	})
}

// Patch User to data
// @Summary Patch User
// @Description Patch User
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param User body UserPost true "Patch User"
// @Param id path int true "User ID"
// @Success 200 {object} common.ResponseHTTP{data=UserPost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /api/users/{id} [patch]
func PatchUsers(contx *fiber.Ctx) error {
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
	patch_User := new(UserPost)
	if err := contx.BodyParser(&patch_User); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// then validating
	if err := validate.Struct(patch_User); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	patch_User.Password = utils.HashFunc(patch_User.Password)
	// startng update transaction
	User := new(models.User)
	tx := db.Begin()
	if err := db.Model(&User).Where("id = ?", id).First(&User).UpdateColumns(*patch_User).Error; err != nil {
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
		Message: "Success Updating a User.",
		Data:    User,
	})
}

// DeleteUsers function removes a User by ID
// @Summary Remove User by ID
// @Description Remove User by ID
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/users/{id} [delete]
func DeleteUsers(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var User models.User
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
	if err := db.Where("id = ?", id).First(&User).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	db.Delete(&User)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a User.",
		Data:    User,
	})
}
