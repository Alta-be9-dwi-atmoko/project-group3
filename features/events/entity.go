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
	DetailEventBusiness(idEvent int) (response Core, err error)
	UpdateEventBusiness(data Core, id int, idUser int) (response int, err error)
	DetailImageEventBusiness(idEvent int) (response string, err error)
}

type Data interface {
	InsertData(data Core) (response int, err error)
	DetailEventData(idEvent int) (response Core, err error)
	UpdateEventData(data Core, id int, idUser int) (response int, err error)
	DetailImageEventData(idEvent int) (response string, err error)
}
