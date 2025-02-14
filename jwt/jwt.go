package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 实际使用时应从配置或环境变量读取
var accessTokenKey = []byte("kfgakgfuagfuhb65441@#$%uihafi")
var refreshTokenKey = []byte("huhaiubuykfgcbcj6548463212!$hgadyu")

type Claims struct {
	ID uint
	jwt.RegisteredClaims
}

// 使用双token
func GetToken(id uint) (string, string, error) {
	// accessToken过期时间一周, refreshToken过期时间一月
	accessTokenTime := time.Now().Add(7 * 24 * time.Hour)
	refreshTokenTime := time.Now().Add(4 * 7 * 24 * time.Hour)

	accessClaims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my",
			Subject:   "token",
		},
	}
	refreshClaims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my",
			Subject:   "token",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenStr, err := accessToken.SignedString(accessTokenKey)
	if err != nil {
		return "", "", err
	}
	refreshTokenStr, err := refreshToken.SignedString(refreshTokenKey)
	if err != nil {
		return "", "", err
	}
	return accessTokenStr, refreshTokenStr, nil
}

func ParseAccessToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return accessTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func ParseRefreshToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
