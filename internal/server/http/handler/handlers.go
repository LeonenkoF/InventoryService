package handler

import (
	service "LeonenkoF/PC_inventory/internal/service/domain"

	"github.com/go-chi/chi/v5"
)

func SetHandlers(router chi.Router, is *service.InventoryService) {
	ir := &inventoryRoutes{
		is: is,
	}

	router.Route("/inventory", func(router chi.Router) {
		router.Get("/", ir.GetAllInventory)
	})

	router.Route("/auth", func(router chi.Router) {
		router.Post("/sign-in", ir.signIn)
	})

	router.Group(func(router chi.Router) {
		router.Use(ir.userIdentity)
		router.Post("/sign-up", ir.signUp)
		router.Post("/inventory", ir.AddNewInventory)
		router.Delete("/inventory", ir.DeleteInventoryHandler)
	})
}
