package responses

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"semay.com/admin/database"
	"semay.com/admin/models"
	"semay.com/common"
	"semay.com/videoutils"
)

type VideoPost struct {
	Video string `json:"video_name"`
}

type PicturePost struct {
	Picture string `json:"picture_name"`
}

// PostVideo is a function to get a Users by ID
// @Summary PostVideo
// @Description Get Users
// @Security ApiKeyAuth
// @Tags Files
// @Accept multipart/form-data
// @Param payload formData VideoPost true "payload"
// @Param video formData file true  "Video"
// @Success 200 {object} common.ResponsePagination{data=[]string}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /blobvideo [post]
func UploadingVideo(contx *fiber.Ctx) error {
	video, verr := contx.FormFile("video")
	if verr != nil {
		// Handle error
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: verr.Error(),
			Data:    nil,
		})
	}

	randomize_name, _ := videoutils.RandomString(7)
	dest := fmt.Sprintf("./Stream/%s", strings.Split(video.Filename, ".")[0]+"_"+randomize_name+".mp4")
	if serr := contx.SaveFile(video, dest); serr != nil {
		fmt.Println(serr)
	}
	videoutils.CreateVideo(strings.Split(video.Filename, ".")[0] + "_" + randomize_name + ".mp4")
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a BlobData.",
		Data:    "Sucess saving data as Blob",
	})
}

// PostPicture is a function to get a Users by ID
// @Summary PostPicture
// @Description Get Users
// @Security ApiKeyAuth
// @Tags Files
// @Accept multipart/form-data
// @Param payload formData PicturePost true "payload"
// @Param picture formData file true  "Picture"
// @Success 200 {object} common.ResponsePagination{data=[]string}
// @Failure 404 {object} common.ResponseHTTP{}
// @Router /blobpic [post]
func UploadingPicture(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	// fmt.Println(contx.FormFile("picture"))
	picture, err := contx.FormFile("picture")
	if err != nil {
		// Handle error
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	pic_content, _ := picture.Open()
	pic_buf, _ := io.ReadAll(pic_content)

	// saving file
	blob := new(models.BlobPicture)
	blob.BlobPicture = pic_buf
	tx := db.Begin()
	if err := tx.Create(&blob).Error; err != nil {
		tx.Rollback()
		return contx.Status(http.StatusInternalServerError).JSON(common.ResponseHTTP{
			Success: false,
			Message: "Picture Creation Sucess",
			Data:    err,
		})
	}
	tx.Commit()
	// return data if transaction is sucessfull
	return contx.Status(http.StatusOK).JSON(common.ResponseHTTP{
		Success: true,
		Message: "Success register a BlobData.",
		Data:    "Sucess saving data as Blob",
	})
}

// StreamVideo is a function to get a Users by ID
// @Summary StreamVideo
// @Description StreamVideo
// @Tags Files
// @Param name path string true "File Name"
// @Router /stream/{file_name} [get]
func StreamingVideo(contx *fiber.Ctx) error {
	db := database.ReturnSession()
	file_name := contx.Params("file_name")
	var blobs models.BlobVideo
	if res := db.Model(&models.BlobVideo{}).Select("blob_video").Where("name = ?", file_name).Scan(&blobs); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    "nil",
		})
	}
	// contx.Response().Header.Add("Content-Type", "video/mp4")
	contx.Response().Header.Add("Accept", "bytes")

	// return contx.Send(blob_video.BlobVideo)
	contx.Send(blobs.BlobVideo)
	return nil
}

// StreamPicture is a function to get a Users by ID
// @Summary StreamPicture
// @Description StreamPicture
// @Tags Files
// @Router /pics [get]
func StreamingPicture(contx *fiber.Ctx) error {
	db := database.ReturnSession()

	var blobs models.BlobPicture
	if res := db.Model(&models.BlobPicture{}).Select("blob_picture").Where("id = ?", 1).Scan(&blobs); res.Error != nil {
		return contx.Status(http.StatusServiceUnavailable).JSON(common.ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    "nil",
		})
	}
	// return contx.Send(blob_video.BlobVideo)
	return contx.Send(blobs.BlobPicture)
}
