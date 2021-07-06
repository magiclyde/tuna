/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/6 上午10:56
 * @note:
 */

package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(key string, mp map[string]interface{}) (string, error) {
	clam := make(jwt.MapClaims)
	for k, v := range mp {
		clam[k] = v
	}
	clam["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clam)
	return token.SignedString([]byte(key))
}

func ParseToken(key string, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, fmt.Errorf("That's not even a token: %s \n", tokenString)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, fmt.Errorf("Token is either expired or not active yet: %s \n", tokenString)
			} else {
				// Couldn't handle this token
				return nil, fmt.Errorf("Couldn't handle this token: %s \n", tokenString)
			}
		} else {
			// Couldn't handle this token
			return nil, fmt.Errorf("Couldn't handle this token: %s \n", tokenString)
		}
	}
	if !token.Valid {
		return nil, fmt.Errorf("Token invalid: %s \n", tokenString)
	}
	return token, nil
}

func RefreshToken(key string, tokenString string) (string, error) {
	token, err := ParseToken(key, tokenString)
	if err != nil {
		return "", err
	}
	clam, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("RefreshToken ERROR")
	}
	mp := make(map[string]interface{})
	for k, v := range clam {
		mp[k] = v
	}
	return CreateToken(key, mp)
}
