package reports

import (
	"time"
)

type Report struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//UpdatedAt   sql.NullTime `json:"updated_at"`
