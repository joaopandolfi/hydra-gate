package security

import (
	"fmt"
	"time"

	"hydra_gate/config"

	jwt "github.com/golang-jwt/jwt/v4"
)

// CheckJwtToken - Check sended token
func CheckJwtToken(tokenString string) (Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().Server.Security.JWTSecret), nil
	})
	if err != nil {
		return Token{Authorized: false}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return Token{Authorized: false}, fmt.Errorf("invalid Token")
	}

	exps := claims["exp"].(float64)
	if int64(exps) < time.Now().Unix() {
		return Token{Authorized: false}, fmt.Errorf("expired token")
	}

	return Token{
		Authorized: true,
		ID:         claims["id"].(string),
	}, nil
}

// NewJwtToken - Crete token with expiration time
func NewJwtToken(t Token, expMinutes int) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = t.ID
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.Get().Server.Security.JWTSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
