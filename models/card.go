package models

import "time"

type CardModel struct {
	Id           int
	UserId       int
	Title        string
	Balance      int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transactions []TransactionModel
}
