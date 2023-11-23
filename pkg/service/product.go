package service

import (
	"github.com/GRRashid/lamoda"
	"github.com/GRRashid/lamoda/pkg/repository"
)

type ProductService struct {
	products repository.Product
}

func NewProductService(products repository.Product) *ProductService {
	return &ProductService{products: products}
}

func (p *ProductService) Create(product lamoda.RawProduct) (int, error) {
	return p.products.Create(product)
}

func (p *ProductService) GetLast(storageId int) ([]lamoda.Product, error) {
	return p.products.GetLast(storageId)
}

func (p *ProductService) ReservedProduct(ids []int) error {
	return p.products.ReservedProduct(ids)
}

func (p *ProductService) UnreservedProduct(ids []int) error {
	return p.products.UnreservedProduct(ids)
}
