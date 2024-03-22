package helpers

import (
	"github.com/dgrijalva/jwt-go"
	
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}