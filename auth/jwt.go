package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

func GetTokenFromCoockies() {

}

func GetTokenFromHeader(headers []string) (string, error) {
	return "", nil
}

func GetClaims() (jwt.Claims, error) {
	return nil, nil
}

func AddToContext(ctx context.Context, token string) context.Context {
	return nil
}

func GetFromContext(ctx context.Context) (string, error) {
	return "", nil
}
