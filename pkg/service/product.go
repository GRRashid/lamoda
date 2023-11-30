package service

import (
	"fmt"
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
	_, err := p.products.FindStorage(product.StorageId)
	if err != nil {
		return 0, err
	} else {
		return p.products.Create(product)
	}
}

func (p *ProductService) GetLast(storageId int) ([]lamoda.Product, error) {
	return p.products.FindAvailableProducts(storageId)
}

func (p *ProductService) ReservedProduct(ids []int) error {
	updatableProducts, err := p.products.FindUpdatableProducts(ids, "available")
	if err != nil {
		return err
	}

	if len(updatableProducts) == 0 {
		return fmt.Errorf("No products with available status found for update")
	} else {
		return p.products.ReservedProduct(updatableProducts)
	}
}

func (p *ProductService) UnreservedProduct(ids []int) error {
	updatableProducts, err := p.products.FindUpdatableProducts(ids, "reserved")
	if err != nil {
		return err
	}

	if len(updatableProducts) == 0 {
		return fmt.Errorf("No products with available status found for update")
	} else {
		return p.products.UnreservedProduct(updatableProducts)
	}
}
