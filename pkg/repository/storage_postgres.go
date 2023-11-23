package repository

import (
	"fmt"

	"github.com/GRRashid/lamoda"
	"github.com/jmoiron/sqlx"
)

type StoragePostgres struct {
	db *sqlx.DB
}

func NewStoragePostgres(db *sqlx.DB) *StoragePostgres {
	return &StoragePostgres{db: db}
}

func (p *StoragePostgres) CreateStorage(storage lamoda.RawStorage) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (name, accessibility) VALUES ($1, $2) RETURNING id", storagesTable)
	row := tx.QueryRow(createProductQuery, storage.Name, storage.Accessibility)
	err = row.Scan(&id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, tx.Commit()
}
