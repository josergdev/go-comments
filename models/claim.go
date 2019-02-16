package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim model for jwt
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
