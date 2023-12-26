package utility

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// type DeToken struct{
// 	id
// }


func CreateToken(id string,secret string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id" : id,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil{
		return "", errors.New(err.Error())
	}

	return tokenString, nil
}


func DecodeToken(cookie string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok{
			return nil, errors.New("method signature not allowed")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token tidak valid")
	}

	claims := token.Claims.(jwt.MapClaims)
	return &claims, nil
}
