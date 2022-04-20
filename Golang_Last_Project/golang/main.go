package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	helper "github.com/FurkanSamaraz/IsEmpty"
	emailControl "github.com/FurkanSamaraz/emailControl"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
)

var uname, pwd, email string

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "172754"
	dbname   = "postgres"
)

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

func openConnention() *sql.DB {

	psq := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psq)

	if err != nil {
		helper.IsEmpty(err.Error())
	}
	err = db.Ping()
	if err != nil {
		helper.IsEmpty(err.Error())
	}

	return db
}

var userr Names
var tokenString string

func login(w http.ResponseWriter, r *http.Request) {
	var jwtKey = []byte("my_secret_key") //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err.Error())
	}
	fmt.Printf("\n")
	expirationTime := time.Now().Add(5 * time.Minute)
	clamis := &Claims{
		Name: userr.UserModel.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis) //token şifreleme
	tokenString, _ = Token.SignedString(jwtKey)                //şifrelenen tokeni anahtarımıza göndererek imzalı tokeni elde etme

	var people []Names
	db := openConnention()
	r.ParseForm()
	uname := r.FormValue("username")
	pwd := r.FormValue("password")
	email := r.FormValue("email")
	userr.UserModel.Token = tokenString
	userr.Message = "Status OK"
	userr.Successful = "Successful"
	userr.UserModel.Token = tokenString
	fmt.Println(tokenString)
	rows, _ := db.Query("SELECT * FROM userMod")
	for rows.Next() {
		rows.Scan(&userr.UserModel.Id, &userr.UserModel.Username, &userr.UserModel.Password, &userr.UserModel.Email)
		people = append(people, userr)

	}
	if uname == userr.UserModel.Username && pwd == userr.UserModel.Password && email == userr.UserModel.Email {

		peopleByte, _ := json.MarshalIndent(userr, "", "\t")
		w.Write(peopleByte)

	}

	defer db.Close()
}

func bosSignup(w http.ResponseWriter, r *http.Request) {
	unameCheck := helper.IsEmpty(uname)
	pwdCheck := helper.IsEmpty(pwd)
	mailCheck := helper.IsEmpty(email)

	if unameCheck || pwdCheck || mailCheck {

	} else {
		fmt.Fprintf(w, "Error Empty! \n")
	}
}

func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
func register(w http.ResponseWriter, r *http.Request) {
	db := openConnention()
	r.ParseForm()
	var lg LoginModel

	lg.Username = r.FormValue("username")
	lg.Password = r.FormValue("password")
	lg.Email = r.FormValue("email")

	rows, _ := db.Query("SELECT * FROM userMod")
	for rows.Next() {
		rows.Scan(&userr.UserModel.Id, &userr.UserModel.Username, &userr.UserModel.Password, &userr.UserModel.Email)

	}

	uCheck := strings.Contains(lg.Password, lg.Username)
	eCheck := strings.Contains(lg.Password, lg.Email)

	switch {
	case uCheck == true || eCheck == true:
		fmt.Fprintf(w, "Password must not contain username or email.")
	case isValid(lg.Password) != true:
		fmt.Fprintf(w, "Use special characters, numbers, upper and lower case letters in the password.")
	case lg.Username == "" || lg.Password == "" || lg.Email == "":
		fmt.Fprintf(w, "cannot be empty")
	case userr.UserModel.Username == lg.Username:
		fmt.Fprintf(w, "username is used")
	case emailControl.CheckEmail(lg.Email) == true:

		db.Exec("INSERT INTO userMod(username,password,email) VALUES($1,$2,$3)", lg.Username, lg.Password, lg.Email)

		peopleByte, _ := json.MarshalIndent(lg, "", "\t")

		w.Header().Set("Content-Type", "application/json")

		w.Write(peopleByte)

		defer db.Close()
		bosSignup(w, r)
		db.Close()
	default:
		fmt.Fprintln(w, "record failed error email!! ", uname)
	}

}
func products(w http.ResponseWriter, r *http.Request) {

}
func update(w http.ResponseWriter, r *http.Request) {

	db := openConnention()
	r.ParseForm()

	userr.UserModel.Username = r.FormValue("username")
	userr.UserModel.Password = r.FormValue("password")
	userr.UserModel.Email = r.FormValue("email")
	db.Exec("UPDATE userToken SET username=$1,password=$2,email=$3 WHERE id=$4 ", userr.UserModel.Username, userr.UserModel.Password, userr.UserModel.Email, userr.UserModel.Id)

	peopleByte, _ := json.MarshalIndent(userr, "", "\t")

	w.Header().Set("Content-Type", "application/json")

	w.Write(peopleByte)

	defer db.Close()
	bosSignup(w, r)
	db.Close()
}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/update", update)
	http.ListenAndServe(":8080", mux)
}
