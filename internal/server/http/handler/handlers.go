package handler

import (
	service "LeonenkoF/PC_inventory/internal/service/domain"

	"github.com/go-chi/chi/v5"
)

func SetHandlers(router chi.Router, is *service.InventoryService) {
	ir := &inventoryRoutes{
		is: is,
	}

	router.Get("/inventory", ir.GetAllInventory)
}
