package storage

import (
	entity "LeonenkoF/PC_inventory/internal/entities/PC_inventory"
	"LeonenkoF/PC_inventory/pkg/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewStorage(cfg *config.Db_connect) (*Postgres, error) {
	connStr := "postgres://postgres:qwerty@localhost:5432/PC_inventory?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed connect to database: %v", err)
	}
	defer db.Close()

	return &Postgres{db: db}, nil

}

func (db *Postgres) GetInventory() ([]entity.Inventory, error) {

	stmt, err := db.db.Query(`SELECT department.dep_name,inventory.*
	FROM inventory
	JOIN department ON department.dep_id = inventory.fk_dep_id;`)

	if err != nil {
		log.Printf("Databese query error: %v", err)
		return nil, err
	}

	defer stmt.Close()

	inventory := []entity.Inventory{}

	for stmt.Next() {
		i := entity.Inventory{}
		err := stmt.Scan(&i.Dep_name, &i.Inventory_id, &i.Fk_dep_id, &i.Full_name, &i.Pc_id, &i.Inventory_num, &i.Invent_monitors, &i.Invent_printer, &i.Invent_mfu, &i.Invent_laptop, &i.Invent_other)

		if err != nil {
			log.Printf("Databese query error: %v", err)
			return nil, err
		}
		inventory = append(inventory, i)
	}
	return inventory, nil
}
