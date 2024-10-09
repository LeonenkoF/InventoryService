package service

import (
	entity "LeonenkoF/PC_inventory/internal/entities/PC_inventory"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt      = "ofgkmogmiamfgkdfgadfg"
	tokenTTL  = time.Hour * 12
	singinKey = "adfhkglsbjjngjlskn"
)

func (is *InventoryService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return is.repo.CreateUser(user)
}

func (is *InventoryService) GenerateToken(user entity.User) (string, error) {
	fmt.Println(user.Username, generatePasswordHash(user.Password))
	user, err := is.repo.GetUser(user.Username, generatePasswordHash(user.Password))

	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.User_id
	claims["username"] = user.Username
	claims["exp"] = tokenTTL

	tokenString, err := token.SignedString([]byte(singinKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
