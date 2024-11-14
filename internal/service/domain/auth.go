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
	signinKey = "adfhkglsbjjngjlskn"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(user entity.User) (string, error)
	ParseToken(accessToken string) (int, error)
}

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.User_id,
	})

	tokenString, err := token.SignedString([]byte(signinKey))

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

func (is *InventoryService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(signinKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
