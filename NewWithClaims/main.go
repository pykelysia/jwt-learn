package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	mapClaims := jwt.MapClaims{
		"iss": "pyke",                // 发行者
		"sub": "pykelysia.github.io", // 接收者
		"aud": "elysia",              // 接收者
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	fmt.Println(token)
}
