package main

import (
	"crypto/rand"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRandomKey() ([]byte, error) {

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func ParseJWTWithClaims(tokenString string, jwtkey []byte, options ...jwt.ParserOption) (jwt.Claims, error) {
	mc := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, mc,
		func(token *jwt.Token) (interface{}, error) {
			return jwtkey, nil
		}, options...)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf(("invalid token"))
	}
	return mc, nil
}

func main() {

	jwtkey, err := GenerateRandomKey()
	if err != nil {
		panic(err)
	}

	mapClaims := jwt.MapClaims{
		"iss": "pyke",                // 发行者
		"sub": "pykelysia.github.io", // 接收者
		"aud": "elysia",              // 接收者
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		panic(err)
	}

	claims, err := ParseJWTWithClaims(tokenString, jwtkey)
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString, "\n\n", claims)
}
