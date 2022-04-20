package middlewarre

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"main.go/checkError"
)

type Claims struct { //Payload
	Name string `json:"name"`
	jwt.StandardClaims
}

var s string

func JwtMidd(s string) string {

	var jwtKey = []byte("my_secret_key") //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
	_, err := rand.Read(jwtKey)
	checkError.ErrorContr(err, "1")

	fmt.Printf("\n")
	expirationTime := time.Now().Add(15 * time.Minute)
	clamis := &Claims{
		Name: s,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis) //token şifreleme
	tokenString, _ := Token.SignedString(jwtKey)               //şifrelenen tokeni anahtarımıza göndererek imzalı tokeni elde etme
	return tokenString
}
