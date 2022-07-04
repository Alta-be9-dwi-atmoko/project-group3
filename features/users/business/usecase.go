package business

import (
	"errors"
	"project/group3/features/users"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBusiness(dataUser users.Data) users.Business {
	return &userUseCase{
		userData: dataUser,
	}
}

func (uc *userUseCase) CreateData(input users.Core) (row int, err error) {
	if input.Name == "" || input.Email == "" || input.Password == "" || input.AvatarUrl == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = uc.userData.InsertData(input)
	return row, err
}
