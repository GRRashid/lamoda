package repository

import (
	"github.com/GRRashid/lamoda"
	"github.com/jmoiron/sqlx"
)

type Product interface {
	Create(input lamoda.RawProduct) (int, error)
	GetLast(storageId int) ([]lamoda.Product, error)
	ReservedProduct(ids []int) error
	UnreservedProduct(ids []int) error
	CountProductsInStorage(storageId int) (int, error)
}

type Storage interface {
	CreateStorage(input lamoda.RawStorage) (int, error)
}

type Repository struct {
	Product
	Storage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductPostgres(db),
		Storage: NewStoragePostgres(db),
	}
}
