package repository

import "github.com/Rest_Savings/pkg/database"

type Repository struct {
	Gormdb *database.GormDB
}

func New(db *database.GormDB) *Repository {
	return &Repository{
		Gormdb: db,
	}
}
