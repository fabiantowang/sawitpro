package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrClaimsInvalid = errors.New("provided claims do not match the schema")

type JwtCustomClaims struct {
	Userid string `json:"userid"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	key, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))

	// Generate encoded token and return it.
	return token.SignedString(key)
}

func GetClaims(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		key, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))

		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return claims, ErrClaimsInvalid
	}

	return claims, nil
}

func GetClaimsCtx(ctx context.Context) (*JwtCustomClaims, error) {
	claims, ok := ctx.Value(JWTClaimsContextKey).(*JwtCustomClaims)
	if !ok {
		return nil, ErrClaimsInvalid
	}

	return claims, nil
}
