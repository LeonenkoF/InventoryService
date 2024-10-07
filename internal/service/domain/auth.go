package service

import (
	entity "LeonenkoF/PC_inventory/internal/entities/PC_inventory"
	"crypto/sha1"
	"fmt"
)

const (
	salt = "ofgkmogmiamfgkdfgadfg"
)

func (is *InventoryService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return is.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
