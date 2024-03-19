package entity

import "github.com/google/uuid"

type Actor struct {
	Id        uuid.UUID `db:"id"`
	Sex       byte      `db:"sex"`
	BithDate  string    `db:"bith_date"`
	IsDeleted bool      `db:"is_deleted"`
	CreatedAt string    `db:"created_at"`
	UpdatedAt string    `db:"updated_at"`
	DeletedAt string    `db:"deleted_at"`
}
