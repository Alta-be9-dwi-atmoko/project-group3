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
}

type Data interface {
}
