package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtSecret = []byte("NzZbFMr2B+3j7BZvin8BCIEr/JcSPTdBvmO0MLjKDDE=")

func GenerateJWT(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    return token, err
}
