package handler

import (
	service "LeonenkoF/PC_inventory/internal/service/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type inventoryRoutes struct {
	is *service.InventoryService
}

func (ir *inventoryRoutes) GetAllInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	data, err := ir.is.Inventory()

	if err != nil {
		log.Printf("InventoryRoutes - Inventory - ir.is.Inventory: %v", err)
	}

	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusBadRequest)
		fmt.Printf("err")
		return
	}
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		log.Printf("error while writing data ResponseWriter: %v", err)
	}
}
