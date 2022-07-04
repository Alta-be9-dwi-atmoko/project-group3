package factory

import (
	_userBusiness "project/group3/features/users/business"
	_userData "project/group3/features/users/data"
	_userPresentation "project/group3/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// dbConn := config.InitDB()
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		UserPresenter: userPresentation,
	}
}
