package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtClaims struct {
	jwt.StandardClaims
}

var _secret []byte
var _issuer string
var _expireDays int

// InitJwt 初始化jwt
func InitJwt(secret string, issuer string, expireDays int) {
	_secret = []byte(secret)
	_issuer = issuer
	_expireDays = expireDays
}

// GenerateJwtToken 生成JwtToken
func GenerateJwtToken(claims *JwtClaims) (string, error) {
	claims.Issuer = _issuer
	claims.NotBefore = int64(time.Now().Unix())
	claims.ExpiresAt = int64(time.Now().AddDate(0, 0, _expireDays).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(_secret)
}

// ParseJwtToken 解码JwtToken
func ParseJwtToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return _secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
