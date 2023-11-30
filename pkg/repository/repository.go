package repository

import (
	"github.com/GRRashid/lamoda"
	"github.com/jmoiron/sqlx"
)

type Product interface {
	Create(input lamoda.RawProduct) (int, error)
	FindAvailableProducts(storageId int) ([]lamoda.Product, error)
	ReservedProduct(updatableProducts []int) error
	UnreservedProduct(updatableProducts []int) error
	FindStorage(storageId int) (int, error)
	FindUpdatableProducts(ids []int, status string) ([]int, error)
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
