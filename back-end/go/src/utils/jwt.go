package util

import (
	// "encoding/json"

	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func EncodeJwtToken(data map[string]interface{}) (string, error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")
	data["expiredAt"] = time.Now().Add(time.Hour * 1).Unix()
	hmacSampleSecret := []byte(SECRET_KEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))
	return token.SignedString(hmacSampleSecret)

}

func DecodeJwtToken(encodedToken string) (map[string]interface{}, error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")
	hmacSampleSecret := []byte(SECRET_KEY)

	token, parse_error := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if parse_error != nil {
		return nil, fmt.Errorf("erro ao parsear o jwt %v", parse_error)
	}

	claims, isClaimsValid := token.Claims.(jwt.MapClaims)
	if isClaimsValid && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("claim, token, hmacSampleSecret is invalid")
	}

}
