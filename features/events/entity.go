package events

import (
	"time"
)

type Core struct {
	ID          int
	User        User
	Image       string
	EventName   string
	Category    string
	Link        string
	Lat         string
	Long        string
	Quota       uint
	Date        string
	Time        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	CreateEventBusiness(data Core) (response int, err error)
}

type Data interface {
	InsertData(data Core) (response int, err error)
}
