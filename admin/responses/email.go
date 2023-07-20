package responses

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"semay.com/bluerabbit"
	"semay.com/common"
)

type EmailMessage struct {
	Emails  []string `json:"emails" validate:"required"`
	Subject string   `json:"subject" validate:"required"`
	Message string   `json:"message" validate:"required"`
}

// Send Email to list of users using rabbit
// @Summary Send Email to
// @Description Sending Email
// @Tags Utilities
// @Accept json
// @Produce json
// @Param User body EmailMessage true "messages"
// @Success 200 {object} common.ResponseHTTP{data=EmailMessage}
// @Failure 404 {object} common.ResponseHTTP{}
// @Failure 503 {object} common.ResponseHTTP{}
// @Router /api/email [post]
func SendEmail(contx *fiber.Ctx) error {
	validate := validator.New()
	//   connection and channels from rabbitmq
	connection, channel := bluerabbit.BrokerConnect()

	defer connection.Close()
	defer channel.Close()
	//validating post data
	posted_message := new(EmailMessage)

	//first parse post data
	if err := contx.BodyParser(&posted_message); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_message); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// make message publishable
	// Create a message to publish.
	ser_message, _ := json.Marshal(posted_message)
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(ser_message),
		Type:        "BULK_MAIL",
	}

	//send to rabbit app module qeue using channel
	// Attempt to publish a message to the queue.
	if err := channel.Publish(
		"",          // exchange
		"blueadmin", // queue name
		false,       // mandatory
		false,       // immediate
		message,     // message to publish
	); err != nil {
		return contx.Status(http.StatusBadRequest).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	// close connection and channel of the rabbitmq server

	return contx.JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success get all Users.",
		Data:    posted_message,
	})

}
