package auth

import "github.com/golang-jwt/jwt"

type Service interface {
	GenerateToken(userID int) (string, error)
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
