package domain

import (
	"github.com/form3tech-oss/jwt-go"
)

type UserJwt struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
