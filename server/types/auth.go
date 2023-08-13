package types

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	CustomClaims struct {
		UserId   string `json:"user_id"`
		IsActive bool   `json:"is_active"`
		Email    string `json:"email"`
	}
	AuthCustomClaims struct {
		jwt.StandardClaims
		CustomClaims
	}
	LoginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RegisterBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name string `json:"name"`
		
	}
)
