package responses

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
	"semay.com/utils"
)

type UserGet struct {
	ID             uint          `json:"id,omitempty"`
	UUID           uuid.UUID     `json:"uuid,omitempty"`
	Email          string        `json:"email,omitempty"`
	DateRegistered time.Time     `json:"date_registered,omitempty"`
	Disabled       bool          `json:"disabled"`
	Roles          []models.Role `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

type UserPost struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type UserPatch struct {
	Email    string `json:"email" example:"someone@domain.com"`
	Disabled bool   `json:"disabled" example:"true"`
}

type UserPassword struct {
	Email    string `validate:"required" json:"email" example:"someone@domain.com"`
	Password string `validate:"required" json:"password"`
}

// GetUsers is a function to get a Users by ID
// @Summary Get Users
// @Description Get Users
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param size query int true "page size"
// @Success 200 {object} common.ResponsePagination{data=[]UserGet}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /users [get]
func GetUsers(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var users []UserGet
	Page, _ := strconv.Atoi(contx.Query("page"))
	Limit, _ := strconv.Atoi(contx.Query("size"))
	if Page == 0 || Limit == 0 {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Not Allowed, Bad request",
			Data:    nil,
		})
	}

	result, err := common.PaginationPureModel(db, models.User{}, users, uint(Page), uint(Limit))
	if err != nil {
		return contx.JSON(common.ResponseHTTP{
			Success: true,
			Message: "Success get all Users.",
			Data:    "something",
		})
	}

	return contx.Status(http.StatusOK).JSON(result)

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
// @Router /users/{id} [get]
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
	var user_get UserGet
	if res := db.Model(&models.User{}).Preload(clause.Associations).Where("id = ?", id).First(&users); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	mapstructure.Decode(users, &user_get)
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one User.",
		Data:    &user_get,
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
// @Router /userrole/{user_id} [get]
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

	var roles_get []RoleGet
	var roles []models.Role
	var user models.User
	if res := db.Model(&models.User{ID: uint(user_id)}).Find(&user); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	db.Model(&user).Association("Roles").Find(&roles)
	mapstructure.Decode(roles, &roles_get)
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success got one route.",
		Data:    &roles_get,
	})
}

// Add Users
// @Summary Add a new Users
// @Description Add Users
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param User body UserPost true "Add User"
// @Success 200 {object} common.ResponseHTTP{data=UserPost}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /users [post]
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
	if err := tx.Create(&User).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "User Creation Failed",
			Data:    err,
		})
	}
	tx.Commit()
	user := UserGet{}

	mapstructure.Decode(User, &user)
	// return data if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a User.",
		Data:    user,
	})
}

// Update User Details
// @Summary Patch User
// @Description Patch User
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param User body UserPost true "Patch User"
// @Param id path int true "User ID"
// @Success 200 {object} common.ResponseHTTP{data=UserPatch}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /users/{id} [patch]
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
	patch_User := new(UserPatch)
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
	// patch_User.Password = utils.HashFunc(patch_User.Password)
	// startng update transaction
	User := new(models.User)
	var user UserGet
	User.ID = uint(id)
	tx := db.Begin()
	if err := db.Model(&User).UpdateColumns(*patch_User).Update("disabled", patch_User.Disabled).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	tx.Commit()

	mapstructure.Decode(User, &user)
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Updating a User.",
		Data:    user,
	})
}

// Activate/Deactivate User
// @Summary Activate/Deactivate User
// @Description Activate/Deactivate User
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param status query bool true "Disabled"
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /users/{user_id} [put]
func ActivateDeactivateUser(contx *fiber.Ctx) error {
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

	// Getting Query Parameter
	status := contx.QueryBool("status")

	// Fetching User
	var user models.User
	user.ID = uint(user_id)

	//Updating Didabled Status
	tx := db.Begin()
	if err := db.Model(&user).Update("disabled", status).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err.Error(),
		})
	}
	tx.Commit()
	var response_user UserGet
	mapstructure.Decode(user, &response_user)
	response_user.Disabled = status
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Updating a User.",
		Data:    response_user,
	})
}

// Update User Password Details
// @Summary Put User
// @Description Put User
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body UserPassword true "Password User"
// @Param reset query bool true "Reset Password"
// @Success 200 {object} common.ResponseHTTP{data=UserGet}
// @Failure 400 {object} common.ResponseHTTP{}
// @Router /users 	[put]
func ChangePassword(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()
	// get query parms
	reset_password := contx.QueryBool("reset")

	// first parsing
	patch_User := new(UserPassword)
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

	// startng update transaction
	var user_q models.User
	if err := db.Model(&user_q).Where("email =?", patch_User.Email).Find(&user_q).Error; err != nil {

		return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}

	var user UserGet

	if !reset_password {
		tx := db.Begin()
		patch_User.Password = utils.HashFunc(patch_User.Password)
		if err := db.Model(&user_q).UpdateColumns(*patch_User).Error; err != nil {
			tx.Rollback()
			return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
				Success: false,
				Message: "Record not Found",
				Data:    err,
			})
		}
		tx.Commit()
	} else {
		tx := db.Begin()
		patch_User.Password = utils.HashFunc("default@123")
		if err := db.Model(&user_q).UpdateColumns(*patch_User).Error; err != nil {
			tx.Rollback()
			return contx.Status(http.StatusNotFound).JSON(common.ResponseHTTP{
				Success: false,
				Message: "Record not Found",
				Data:    err,
			})
		}
		tx.Commit()
	}

	mapstructure.Decode(user_q, &user)
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Updating a Password.",
		Data:    user,
	})
}

// DeleteUsers function removes a User by ID
// @Summary Remove User by ID
// @Description Remove User by ID
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} common.ResponseHTTP{}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /users/{id} [delete]
func DeleteUsers(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	var user models.User
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
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Record not Found",
			Data:    err,
		})
	}
	if user.Email == "superuser@mail.com" || user.Email == "standarduser@mail.com" || user.Email == "adminuser@mail.com" {
		return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
			Success: true,
			Message: "You are not allowed to Delete this specific User.",
			Data:    user,
		})
	}
	db.Delete(&user)
	tx.Commit()
	// return value if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success Delete a User.",
		Data:    user,
	})
}
