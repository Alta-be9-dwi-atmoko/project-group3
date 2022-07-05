package data

import (
	"errors"
	"project/group3/features/comments"

	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		DB: db,
	}
}

func (repo *mysqlCommentRepository) InsertData(input comments.Core) (row int, err error) {
	comment := FromCore(input)
	resultCreate := repo.DB.Create(&comment)
	if resultCreate.Error != nil {
		return 0, resultCreate.Error
	}
	if resultCreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert data, your email is already registered")
	}
	return int(resultCreate.RowsAffected), nil
}
