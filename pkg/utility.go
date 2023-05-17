package pkg

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)

		claims["email"] = email
		claims["exp"] = time.Now().Add(time.Minute * 1).Unix() // token akan expired dalam 1 menit

		secretKey := []byte(GetEnv("SECRET_KEY"))

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return "", errors.New("failed to generate token")
		}

		return tokenString, nil
}