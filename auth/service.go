package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

// When a production mode, secret key must be in .env file
var SECRET_KEY = []byte("CROWDFUNDING_s3cr3t_k3y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// Create claim for payload token
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Masukkan token dengan secret key
	signedToken, err := token.SignedString(SECRET_KEY)
	// Check if error
	if err != nil {
		return signedToken, err
	}

	// If no error, return signedToken
	return signedToken, nil
}
