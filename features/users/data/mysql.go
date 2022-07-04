package data

import (
	"errors"
	"fmt"
	"project/group3/features/users"

	_bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.Data {
	return &mysqlUserRepository{
		DB: db,
	}
}

func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	user := FromCore(input)
	passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorHash != nil {
		fmt.Println("Error hash", errorHash.Error())
	}
	user.Password = string(passwordHashed)
	resultcreate := repo.DB.Create(&user)
	if resultcreate.Error != nil {
		return 0, resultcreate.Error
	}
	if resultcreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert data, your email is already registered")
	}
	return int(resultcreate.RowsAffected), nil
}
