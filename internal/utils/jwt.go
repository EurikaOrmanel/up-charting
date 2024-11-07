package utils

import (
	"os"
	"time"

	"crypto/rand"
	"encoding/hex"
	"fmt"

	"EurikaOrmanel/up-charter/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateAccessToken(user models.Admin, expiresIn int) (string, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(expiresIn))
	claims := jwt.MapClaims{
		"id":       user.ID,
		"verified": user.Verified,
		"exp":      expiresAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSecret := os.Getenv("TOKEN_SECRET")

	t, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return "Bearer " + t, nil

}

func GenerateRefreshToken() (string, error) {
	uuid := uuid.New().String()
	randomBase, err := RandomHex(34)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", uuid, randomBase), nil
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func VerifyJWT(tokenString string) (bool, error) {
	tokenSecret := os.Getenv("TOKEN_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
