package presentation

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"project/group3/features/users"
	_requestUser "project/group3/features/users/presentation/request"
	_helper "project/group3/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(userBusiness users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: userBusiness,
	}
}

func (h *UserHandler) PostUser(c echo.Context) error {
	userReq := _requestUser.User{}
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	file, errFile := c.FormFile("product_picture")

	if errFile != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	srcFile, errSrcFile := file.Open()

	if errSrcFile != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	fileByte, _ := ioutil.ReadAll(srcFile)
	fileType := http.DetectContentType(fileByte)
	fileSize := file.Size

	if fileType != "image/png" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	if fileSize < 1024 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	fileName := "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"

	errPermission := ioutil.WriteFile(fileName, fileByte, 0777)

	if errPermission != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	defer srcFile.Close()

	userReq.AvatarUrl = "http://127.0.0.1:80/" + fileName
	dataUser := _requestUser.ToCore(userReq)
	row, errCreate := h.userBusiness.CreateData(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your email is already registered"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
