package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const ACCESS_TOKEN_TIME = 18

func MakeToken(c jwt.Claims, key interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	s, e := token.SignedString(key)

	if e != nil {
		panic(e.Error())
	}

	return s
}

func ParseToken(tokenString string) (interface{}, error) {
	// get PublicKey
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	publicKey, err := LoadRSAPublicKeyFromDisk(dir + "/keys/public.pem")
	if err != nil {
		return "", err
	}
	//
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	if claims, ok := token.Claims.(*jwt.StandardClaims); token.Valid && ok {
		return claims, nil
	}
	return nil, errors.New("Eo biet loi gi")
}

func GenToken(sub string) (string, error) {
	//get private Key
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	signingkey, err := LoadRSAPrivateKeyFromDisk(dir + "/keys/private.pem")
	if err != nil {
		return "", err
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(ACCESS_TOKEN_TIME) * time.Minute).Unix(),
		Issuer:    "Issuer-Test",
		Subject:   sub,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signingkey)
}
