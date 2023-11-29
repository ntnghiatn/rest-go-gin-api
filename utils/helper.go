package utils

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt"
)

func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	// fmt.Println(location)
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

/*
func CreateToken(user string) (string, error) {

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"level1",
		CustomerInfo{user, "human"},
	}

	return t.SignedString(signKey)
}
*/
