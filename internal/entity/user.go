package entity

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `db:"id"`
	Login     string    `db:"login"`
	Password  string    `db:"hashed_password"`
	Role      string    `db:"role"`
	IsDeleted bool      `db:"is_deleted"`
	CreatedAt string    `db:"created_at"`
	UpdatedAt string    `db:"updated_at"`
	DeletedAt string    `db:"deleted_at"`
}
