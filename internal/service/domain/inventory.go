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

func (is *InventoryService) GetInventoryList() ([]entity.Inventory, error) {
	res, err := is.repo.GetInventory()

	if err != nil {
		return nil, fmt.Errorf("InventoryService - GetInventoryList - is.repo.GetinventoryList error: %v", err)
	}
	return res, nil
}

func (is *InventoryService) AddInventory(input entity.Inventory) error {
	err := is.repo.AddInventory(input)

	if err != nil {
		return fmt.Errorf("InventoryService - AddInventory - is.repo.Getinventory error: %v", err)
	}

	if len(input.Full_name) == 0 || len(input.Pc_id) == 0 || len(input.Inventory_num) == 0 || len(input.Invent_monitors) == 0 {
		return fmt.Errorf("full_name or pc_id or Inventory_num or Invent_monitors must be filled out")
	}
	return nil
}
