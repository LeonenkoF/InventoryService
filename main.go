package main

import (
	"LeonenkoF/PC_inventory/internal/server/http/handler"
	service "LeonenkoF/PC_inventory/internal/service/domain"
	storage "LeonenkoF/PC_inventory/internal/storage/postgres/postgresql"
	"LeonenkoF/PC_inventory/pkg/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// @title           Inventory Service API
// @version         1.0
// @description     API server for inventory service.

// @host      localhost:7540
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Println("Error creating config %v", err)
	}

	pg, err := storage.NewStorage(&cfg.Db_connect)
	if err != nil {
		log.Println("error creating database connection: %v", err)
	}

	is := service.NewInventoryService(pg)

	r := chi.NewRouter()

	handler.SetHandlers(r, is)
	r.Handle("/*", http.FileServer(http.Dir("./web")))

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
