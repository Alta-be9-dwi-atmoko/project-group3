package data

import (
	"fmt"
	"project/group3/features/events"

	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	db *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Data {
	return &mysqlEventRepository{
		db: conn,
	}
}

func (repo *mysqlEventRepository) InsertData(input events.Core) (response int, err error) {
	event := fromCore(input)
	result := repo.db.Create(&event)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}

	return int(result.RowsAffected), err
}

func (repo *mysqlEventRepository) DetailEventData(id int) (response events.Core, err error) {
	var dataEvent Event

	result := repo.db.Preload("User").First(&dataEvent, "id = ?", id)

	if result.RowsAffected != 1 {
		return events.Core{}, fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return events.Core{}, result.Error
	}

	return dataEvent.toCore(), nil
}

func (repo *mysqlEventRepository) UpdateEventData(editData events.Core, id int, idUser int) (response int, err error) {

	event := fromCore(editData)
	event_ := fromCore(editData)

	searchEvent := repo.db.First(&event_, "id = ?", id)

	if searchEvent.RowsAffected != 1 {
		return 0, fmt.Errorf("failed update event")
	}

	if event_.UserID != uint(idUser) {
		return 0, fmt.Errorf("failed update event")
	}

	result := repo.db.Model(Event{}).Where("id = ?", id).Updates(&event)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlEventRepository) DetailImageEventData(id int) (response string, err error) {
	var dataEvent Event

	result := repo.db.Preload("User").First(&dataEvent, "id = ?", id)

	if result.RowsAffected != 1 {
		return "", fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return "", result.Error
	}

	return dataEvent.Image, nil
}

func (repo *mysqlEventRepository) DeleteEventData(id int, idUser int) (row int, err error) {
	var dataEvent Event

	searchProduct := repo.db.Find(&dataEvent, id)

	if searchProduct.RowsAffected != 1 {
		return 0, fmt.Errorf("failed delete event")
	}

	if searchProduct.Error != nil {
		return 0, searchProduct.Error
	}

	if dataEvent.UserID != uint(idUser) {
		return 0, fmt.Errorf("failed delete event")
	}

	result := repo.db.Delete(&dataEvent, id)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}
