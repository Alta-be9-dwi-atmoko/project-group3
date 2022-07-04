package routes

import (
	"project/group3/factory"
	_middleware "project/group3/features/middlewares"
	_validatorEvents "project/group3/validator/events"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()

	e.HTTPErrorHandler = _validatorEvents.ErroHandlerEvent

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("/users", presenter.UserPresenter.PostUser)
	e.POST("/login", presenter.UserPresenter.LoginAuth)
	e.PUT("/users/:id", presenter.UserPresenter.PutUser, _middleware.JWTMiddleware())

	// event
	e.POST("/events", presenter.EventPresenter.CreateEvent, _middleware.JWTMiddleware())

	return e

}
