package utils

import (
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(uid int64) (string, error) {
	claims := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
			Issuer:    "yookie",
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenObj.SignedString(viper.GetString("jwtKey"))
}

func ParseToken(token string) (*Claims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return viper.GetString("jwtKey"), nil
	})

	if tokenObj != nil {
		if claims, ok := tokenObj.Claims.(*Claims); ok && tokenObj.Valid {
			return claims, nil
		}
	}
	return nil, err
}
