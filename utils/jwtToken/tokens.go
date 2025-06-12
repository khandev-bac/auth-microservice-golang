package jwttoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/khandev-bac/lemon/config"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

var JWT_A_KEY = config.AppConfig.JwtkeysAccess
var JWT_R_KEY = config.AppConfig.JwtkeysRefresh

func GenerateTokens(id uuid.UUID, email string) (*Tokens, error) {
	AccessTokenclaims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}
	AccessStr := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenclaims)
	AccessToken, err := AccessStr.SignedString([]byte(JWT_A_KEY))
	if err != nil {
		return nil, fmt.Errorf("%s", "failed to create access token"+err.Error())
	}
	RefreshTokenclaims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(90 * 24 * time.Hour).Unix(),
	}
	RefreshStr := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshTokenclaims)
	RefreshToken, err := RefreshStr.SignedString([]byte(JWT_R_KEY))
	if err != nil {
		return nil, fmt.Errorf("%s", "failed to create refresh token"+err.Error())
	}
	return &Tokens{
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
	}, nil
}

func VerifyJWTAccessToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(JWT_A_KEY), nil
	})
	if err != nil {
		return nil, errors.New("error for verifying access token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New("token expired")
		}
	}
	return claims, nil
}
func VerifyJWTRefreshToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(JWT_R_KEY), nil
	})
	if err != nil {
		return nil, errors.New("error for verifying refresh token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New("token expired")
		}
	}
	return claims, nil
}
