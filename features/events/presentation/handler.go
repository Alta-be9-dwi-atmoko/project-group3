package presentation

import (
	"fmt"
	"project/group3/features/events"
	_requestEvent "project/group3/features/events/presentation/request"
	"time"

	// _responseEvent "project/group3/features/events/presentation/response"
	_middlewares "project/group3/features/middlewares"
	_helper "project/group3/helper"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	EventBusiness events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		EventBusiness: business,
	}
}

func (h *EventHandler) CreateEvent(c echo.Context) error {
	// inisialiasi variabel dengan type struct dari request
	var newEvent _requestEvent.Event

	// binding data event
	errBind := c.Bind(&newEvent)

	validate := validator.New()
	if errValidate := validate.Struct(newEvent); errValidate != nil {
		return errValidate
	}

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	// formfile data image event
	fileData, fileInfo, fileErr := c.Request().FormFile("image")

	// return err jika missing file
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
	}

	// cek ekstension file upload
	extension, err_check_extension := _helper.CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
	}

	// check file size
	err_check_size := _helper.CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
	}

	// memberikan nama file
	fileName := time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := _helper.UploadImageToS3(fileName, fileData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
	}

	// ekstrak token
	idToken, _, errToken := _middlewares.ExtractToken(c)

	// return jika errorToken
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	// inissialisasi newEvent.UserId = idToken(userid)
	newEvent.UserId = idToken
	//
	newEvent.Image = url

	dataEvent := _requestEvent.ToCore(newEvent)
	_, err := h.EventBusiness.CreateEventBusiness(dataEvent)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))

	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))

}
