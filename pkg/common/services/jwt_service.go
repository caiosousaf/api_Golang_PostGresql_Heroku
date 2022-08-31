package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRETKEY"),
		issure: os.Getenv("JWT_ISSURE"),
	}
}

type Claim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: s.issure,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken (token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}