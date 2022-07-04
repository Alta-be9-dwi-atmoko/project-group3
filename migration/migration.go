package migration

import (
	_mUsers "project/group3/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})

}
