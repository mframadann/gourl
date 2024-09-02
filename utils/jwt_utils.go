package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	Sub       interface{}
	SecretKey []byte
	Exp       int64
}

func CreateToken(p *TokenPayload) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": p.Sub,
			"iss": "gourl",
			"exp": p.Exp,
			"iat": time.Now().Unix(),
		},
	)

	token, err := claims.SignedString(p.SecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
