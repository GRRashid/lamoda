package service

import (
	"github.com/GRRashid/lamoda"
	"github.com/GRRashid/lamoda/pkg/repository"
)

type StorageService struct {
	storages repository.Storage
}

func NewStorageService(storages repository.Storage) *StorageService {
	return &StorageService{storages: storages}
}

func (p *StorageService) CreateStorage(storage lamoda.RawStorage) (int, error) {
	return p.storages.CreateStorage(storage)
}
