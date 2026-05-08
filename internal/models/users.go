package models

import "time"

type UsersModel struct {
	Id        int
	Name      string
	CreatedAt time.Time
}
