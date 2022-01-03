package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("GOCAMPAIGN_s3cr3t_k3y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	SignedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return SignedToken, err
	}

	return SignedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken String) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
