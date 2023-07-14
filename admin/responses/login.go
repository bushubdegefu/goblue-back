package responses

// https://morkid.github.io/paginate/

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
	"semay.com/utils"
)

type LoginPost struct {
	GrantType string `json:"grant_type" validate:"required"`
	Email     string `json:"email" validate:"email,min=6,max=32"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

// Login is a function to login by EMAIL and ID
// @Summary Auth
// @Description Login
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body LoginPost true "Login"
// @Success 200 {object} common.ResponseHTTP{data=TokenResponse{}}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /login [post]
func PostLogin(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	validate := validator.New()

	//validating post data
	login_request_data := new(LoginPost)

	//first parse post data
	if err := contx.BodyParser(&login_request_data); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(login_request_data); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	switch login_request_data.GrantType {
	case "authorization_code":
		var user models.User
		res := db.Model(&models.User{}).Preload(clause.Associations).Where("email = ?", login_request_data.Email).First(&user)
		if res.Error != nil {
			return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
				Success: false,
				Message: res.Error.Error(),
				Data:    nil,
			})
		} else if utils.PasswordsMatch(user.Password, login_request_data.Password) {
			roles := make([]string, 0, 20)
			for _, value := range user.Roles {

				roles = append(roles, string(value.Name))
			}
			accessString, _ := utils.CreateJWTToken(user.Email, user.UUID.String(), roles, 60)
			refreshString, _ := utils.CreateJWTToken(user.Email, user.UUID.String(), roles, 65)

			data := TokenResponse{
				AccessToken:  accessString,
				RefreshToken: refreshString,
				TokenType:    "Bearer",
			}
			return contx.Status(http.StatusAccepted).JSON(common.ResponseHTTP{
				Success: true,
				Message: "Authorization Granted",
				Data:    data,
			})
		} else {
			return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
				Success: false,
				Message: "Make sure You are Providing the Correct Credentials",
				Data:    "Authenthication Failed",
			})
		}
		// return "something"
	case "refresh_token":

		claims, err := utils.ParseJWTToken(login_request_data.Token)
		email, _ := claims["email"].(string)
		uuid, _ := claims["uuid"].(string)
		roles, _ := claims["roles"].([]string)
		if err == nil {

			accessString, _ := utils.CreateJWTToken(email, uuid, roles, 60)
			refreshString, _ := utils.CreateJWTToken(email, uuid, roles, 65)
			data := TokenResponse{
				AccessToken:  accessString,
				RefreshToken: refreshString,
				TokenType:    "Bearer",
			}
			return contx.Status(http.StatusAccepted).JSON(common.ResponseHTTP{
				Success: true,
				Message: "Authorization Granted",
				Data:    data,
			})
		}

		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Request Type Unknown",
			Data:    "Currently Not Implemented",
		})
	case "token_decode":
		claims, err := utils.ParseJWTToken(login_request_data.Token)

		if err == nil {
			return contx.Status(http.StatusAccepted).JSON(common.ResponseHTTP{
				Success: true,
				Message: "Token decode sucessfull",
				Data:    claims,
			})
		}
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    "Unknown grant type",
		})
	default:
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Request Type Unknown",
			Data:    "Unknown grant type",
		})
	}

}
