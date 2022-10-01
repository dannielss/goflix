package config

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtConfig struct {
	secretKey string
	issuer    string
}

func NewJWTConfig() *jwtConfig {
	return &jwtConfig{
		secretKey: "secret-key",
		issuer:    "goflix-api",
	}
}

type Claim struct {
	Sum int64 `json:"sum"`
	jwt.StandardClaims
}

func (c *jwtConfig) GenerateToken(id int64) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    c.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(c.secretKey))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (c *jwtConfig) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %s", token)
		}

		return []byte(c.secretKey), nil
	})

	return err == nil
}
