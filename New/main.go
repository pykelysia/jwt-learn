package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func main() {
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Println(token)
}
