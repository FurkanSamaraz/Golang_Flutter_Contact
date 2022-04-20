package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"main.go/beeeb"
	"main.go/login"
	"main.go/protected"
	"main.go/register"
	"main.go/update"
)

var uname, pwd, email string

type UserModel struct {
	Id       int
	Username string
	Password string
	Email    string
	Token    string
}
type LoginModel struct {
	Username string
	Password string
	Email    string
}
type Names struct {
	Successful string
	Message    string
	UserModel  UserModel
}
type Claims struct { //Payload
	Name string `json:"name"`
	jwt.StandardClaims
}

func products(w http.ResponseWriter, r *http.Request) {

}

func main() {

	beeeb.BebMessage("🙁", "Data Encrypted\n ‣ register\n ‣ login\n ‣ protected\n")

	mux := http.NewServeMux()
	mux.HandleFunc("/register", register.Register)
	mux.HandleFunc("/login", login.Login)
	mux.HandleFunc("/update", update.Update)
	mux.HandleFunc("/protected", protected.Protected)
	http.ListenAndServe(":8080", mux)
}
