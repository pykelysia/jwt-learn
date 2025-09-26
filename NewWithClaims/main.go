package main

import (
	"crypto/rand"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	jwtkey := make([]byte, 32)
	if _, err := rand.Read(jwtkey); err != nil {
		panic(err)
	}
	mapClaims := jwt.MapClaims{
		"iss": "pyke",                // 发行者
		"sub": "pykelysia.github.io", // 接收者
		"aud": "elysia",              // 接收者
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, e := token.SignedString(jwtkey)
	if e != nil {
		panic(e)
	}
	fmt.Println(tokenString)
}
