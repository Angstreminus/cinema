package entity

import (
	"github.com/google/uuid"
)

type Movie struct {
	Id          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ReleaseDate string    `db:"release_date"`
	Rating      float64   `db:"rating"`
	IsDeleted   bool      `db:"is_deleted"`
	CreatedAt   string    `db:"created_at"`
	UpdatedAt   string    `db:"updated_at"`
	DeletedAt   string    `db:"deleted_at"`
}
