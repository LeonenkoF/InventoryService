package main

import (
	"LeonenkoF/PC_inventory/pkg/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Println("Ошибка в создании кофига: %s", err)
		return
	}

	fmt.Println(cfg.Dbname)

}
