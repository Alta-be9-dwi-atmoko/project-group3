package presentation

import (
	"net/http"
	"project/group3/features/comments"
	_requestComment "project/group3/features/comments/presentation/request"
	_middleware "project/group3/features/middlewares"
	_helper "project/group3/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(commentBusiness comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: commentBusiness,
	}
}

func (h *CommentHandler) PostComment(c echo.Context) error {
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	commentReq := _requestComment.Comment{}
	err := c.Bind(&commentReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	idFromToken, _, _ := _middleware.ExtractToken(c)
	dataComment := _requestComment.ToCore(commentReq)
	dataComment.UserID = idFromToken
	dataComment.EventID = idEvent
	row, errCreate := h.commentBusiness.CreateData(dataComment)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your email is already registered"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
