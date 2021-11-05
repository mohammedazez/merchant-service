package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	err = godotenv.Load()
	key = os.Getenv("SECRET_KEY")
)

type Service interface {
	GenerateTokenUser(UserID string) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateTokenUser(UserID string) (string, error) {
	claim := jwt.MapClaims{
		"user_id": UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}
