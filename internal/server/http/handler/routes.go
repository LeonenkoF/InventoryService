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

// GetAllInventory godoc
// @Summary      GetAllInventory
// @Description  Show inventory list
// @Tags         Inventory
// @Produce      json
// @Router       /inventory [get]
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

// AddNewInventory godoc
// @Summary      AddNewInventory
// @Security	 ApiKeyAuth
// @Description  Adds a new item in inventory
// @Tags         Inventory
// @Accept		 json
// @Produce      json
// @Param        input body entity.Inventory true "item info"
// @Router       /inventory [post]
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

// DeleteInventory godoc
// @Summary      DeleteInventory
// @Security	 ApiKeyAuth
// @Description  Delete item in inventory
// @Tags         Inventory
// @Accept		 json
// @Param        input body entity.Inventory true "item info"
// @Router       /inventory [delete]
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

// SignUp godoc
// @Summary      SignUp
// @Description  Signs up
// @Tags         Auth
// @Accept		 json
// @Param        input body entity.User true "User info"
// @Router       /sign-up [post]
func (ir *inventoryRoutes) signUp(w http.ResponseWriter, r *http.Request) {
	var input entity.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("invalid input body %v", err), http.StatusBadRequest)
		return
	}

	id, err := ir.is.CreateUser(input)

	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(map[string]interface{}{"id": id})

	if err != nil {
		http.Error(w, fmt.Sprintf("JSON encode error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// SignIn godoc
// @Summary      SignIn
// @Description  Signs in
// @Tags         Auth
// @Accept		 json
// @Param        input body entity.User true "SignIn info"
// @Router       /auth/sign-in [post]
func (ir *inventoryRoutes) signIn(w http.ResponseWriter, r *http.Request) {
	var input entity.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("invalid input body %v", err), http.StatusBadRequest)
		return
	}

	token, err := ir.is.GenerateToken(input)

	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(map[string]interface{}{"token": token})

	if err != nil {
		http.Error(w, fmt.Sprintf("JSON encode error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
