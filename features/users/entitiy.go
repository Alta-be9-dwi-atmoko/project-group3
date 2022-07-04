package users

import "time"

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	CreateData(input Core) (row int, err error)
	LoginUser(authData AuthRequestData) (token, name, avatarUrl string, err error)
}

type Data interface {
	InsertData(input Core) (row int, err error)
	LoginUserDB(authData AuthRequestData) (token, name, avatarUrl string, err error)
}
