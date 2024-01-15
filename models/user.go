package models

import (
	"time"
)

type UserModel struct {
	Id        int
	Firstname string
	Lastname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Cards     []CardModel
}
