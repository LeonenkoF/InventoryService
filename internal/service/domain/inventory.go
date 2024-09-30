package service

import (
	entity "LeonenkoF/PC_inventory/internal/entities/PC_inventory"
	storage "LeonenkoF/PC_inventory/internal/storage/postgres/postgresql"
	"fmt"
)

type InventoryService struct {
	repo *storage.Postgres
}

func NewInventoryService(r *storage.Postgres) *InventoryService {
	return &InventoryService{
		repo: r,
	}
}

func (is *InventoryService) Inventory() ([]entity.Inventory, error) {
	res, err := is.repo.GetInventory()

	if err != nil {
		return nil, fmt.Errorf("InventoryService - Inventory - is.repo.Getinventory error: %v", err)
	}
	return res, nil
}
