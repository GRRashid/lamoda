package service

import (
	"github.com/GRRashid/lamoda"
	"github.com/GRRashid/lamoda/pkg/repository"
)

type Product interface {
	Create(input lamoda.RawProduct) (int, error)
	GetLast(storageId int) ([]lamoda.Product, error)
	ReservedProduct(ids []int) error
	UnreservedProduct(ids []int) error
}

type Storage interface {
	CreateStorage(input lamoda.RawStorage) (int, error)
}

type Service struct {
	Product
	Storage
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(repository.Product),
		Storage: NewStorageService(repository.Storage),
	}
}
