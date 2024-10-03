package storage

import (
	entity "LeonenkoF/PC_inventory/internal/entities/PC_inventory"
	"LeonenkoF/PC_inventory/pkg/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewStorage(cfg *config.Db_connect) (*Postgres, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Dbname, cfg.Sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed connect to database: %v", err)
	}

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

func (db *Postgres) AddInventory(input entity.Inventory) error {

	stmt, err := db.db.Query(`INSERT INTO inventory(fk_dep_id,full_name,pc_id,invent_num,invent_monitors,invent_printer,invent_mfu,invent_laptop,invent_other)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`, input.Fk_dep_id, input.Full_name, input.Pc_id, input.Inventory_num, input.Invent_monitors, input.Invent_printer, input.Invent_mfu, input.Invent_laptop, input.Invent_other)

	if err != nil {
		log.Printf("Postgres - AddInventory - Databese query error: %v", err)
		return err
	}

	defer stmt.Close()

	return nil
}

func (db *Postgres) DeleteInventoryPG(id int) error {

	stmt, err := db.db.Query("DELETE FROM Inventory WHERE inventory_id = $1", id)

	if err != nil {
		log.Printf("Postgres - DeleteInventory - Databese query error: %v", err)
		return err
	}

	defer stmt.Close()

	return nil
}
