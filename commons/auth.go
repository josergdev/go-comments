package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/josergdev/go-comments/models"

	"github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("keys/private.rsa")
	if err != nil {
		log.Fatal("Cannot read private key", err)
	}

	publicBytes, err := ioutil.ReadFile("keys/public.rsa")
	if err != nil {
		log.Fatal("Cannot read public key", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Cannot parse privateBytes to privateKey", err)
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Cannot parse publicBytes to PublicKey", err)
	}
}

// GenerateJWT generate token for the user passed to the client
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour * 2).Unix()
			Issuer: "josergdev",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Cannot sign token", err)
	}

	return result
}
