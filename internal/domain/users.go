package domain

import "time"

type Users struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}
