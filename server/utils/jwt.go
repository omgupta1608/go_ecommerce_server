package utils

import (
	"github.com/dgrijalva/jwt-go"
	cfg "github.com/omgupta1608/aftershoot_task/cfg"
	"github.com/omgupta1608/aftershoot_task/types"
)

func GenerateJWTToken(cc types.CustomClaims) (string, error) {
	var signingMethod = jwt.SigningMethodHS256
	var secretKey = cfg.TOKEN_SECRET
	claim := types.AuthCustomClaims{StandardClaims: jwt.StandardClaims{}, CustomClaims: cc}
	token := jwt.NewWithClaims(signingMethod, claim)

	// sign the token using secret key
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
