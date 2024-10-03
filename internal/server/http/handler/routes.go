package handler

import (
	entity "LeonenkoF/PC_inventory/internal/entities/PC_inventory"
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

	data, err := ir.is.GetInventoryList()

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

func (ir *inventoryRoutes) AddNewInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	input := entity.Inventory{}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, `{"error": "ошибка десериализации JSON"}`, http.StatusBadRequest)
		return
	}

	err = ir.is.AddInventory(input)
	if err != nil {
		http.Error(w, fmt.Sprintf("inventoryRoutes - AddNewInventory - ir.is.AddNewInventory: %v", err), http.StatusBadRequest)
		return
	}

}

func (ir *inventoryRoutes) DeleteInventoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	input := entity.Inventory{}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, `{"error": "ошибка десериализации JSON"}`, http.StatusBadRequest)
		return
	}

	err = ir.is.DeleteInventory(input.Inventory_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("inventoryRoutes - AddNewInventory - ir.is.AddNewInventory: %v", err), http.StatusBadRequest)
		return
	}

}
