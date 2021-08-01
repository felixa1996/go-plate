package utils

import (
	"log"
	"os"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/felixa1996/go-plate/domain"
	"github.com/form3tech-oss/jwt-go"
)

type FelJwt struct {
	JwtMiddleware *jwtmiddleware.JWTMiddleware
}

func (j *FelJwt) GetJWTUser(tokenString string) *domain.UserJwt {
	token2 := strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(token2, &domain.UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	claims, ok := token.Claims.(*domain.UserJwt)
	if !ok && !token.Valid {
		log.Fatalln(err)
	}
	return claims
}
