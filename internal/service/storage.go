package service

import "database/sql"

type Word struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Translation string `json:"translation"`
}

type Report struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
