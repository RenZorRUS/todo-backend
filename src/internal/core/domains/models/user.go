package models

import (
	"time"
)

type User struct {
	ID           uint64    `db:"id"            json:"id"`
	Name         string    `db:"name"          json:"name"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Email        string    `db:"email"         json:"email"`
	CreatedAt    time.Time `db:"created_at"    json:"createdAt"`
	UpdatedAt    time.Time `db:"updated_at"    json:"updatedAt"`
}
