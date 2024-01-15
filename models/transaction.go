package models

import "time"

type TransactionModel struct {
	Id              int
	UserId          int
	CardId          int
	Title           string
	Description     string
	TransactionType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Amount          int64
}
