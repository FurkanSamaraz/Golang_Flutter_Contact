package login

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"main.go/beeeb"
	"main.go/checkError"
	"main.go/database"
	"main.go/middlewarre"
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

var Userr Names
var tokenString string
var Control bool
var s = Userr.UserModel.Username

func Login(w http.ResponseWriter, r *http.Request) {
	tokenString = middlewarre.JwtMidd(Userr.UserModel.Username)

	var people []Names
	db := database.OpenConnention()
	r.ParseForm()
	uname := r.FormValue("username")
	pwd := r.FormValue("password")
	email := r.FormValue("email")

	Userr.Message = "Status OK"
	Userr.Successful = "Successful"
	Userr.UserModel.Token = tokenString
	//fmt.Println(tokenString)
	rows, err := db.Query("SELECT * FROM userMod")
	checkError.ErrorContr(err, "3")
	for rows.Next() {
		rows.Scan(&Userr.UserModel.Id, &Userr.UserModel.Username, &Userr.UserModel.Password, &Userr.UserModel.Email)
		people = append(people, Userr)

	}

	if uname == Userr.UserModel.Username && pwd == Userr.UserModel.Password && email == Userr.UserModel.Email {

		peopleByte, _ := json.MarshalIndent(Userr, "", "\t")
		w.Write(peopleByte)
		beeeb.BebMessage("😊", "Login Successful✅\n ‣ protected\n")
		Control = true

	} else {
		Control = false
	}

	defer db.Close()
}
