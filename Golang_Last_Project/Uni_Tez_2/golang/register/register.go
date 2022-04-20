package register

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	helper "github.com/FurkanSamaraz/IsEmpty"
	emailControl "github.com/FurkanSamaraz/emailControl"
	"main.go/beeeb"
	"main.go/checkError"
	"main.go/database"
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

var userr Names
var uname, pwd, email string

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
func bosSignup(w http.ResponseWriter, r *http.Request) {

	unameCheck := helper.IsEmpty(uname)
	pwdCheck := helper.IsEmpty(pwd)
	mailCheck := helper.IsEmpty(email)

	if unameCheck || pwdCheck || mailCheck {

	} else {
		fmt.Fprintf(w, "Error Empty! \n")
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnention()
	r.ParseForm()
	var lg LoginModel

	lg.Username = r.FormValue("username")
	lg.Password = r.FormValue("password")
	lg.Email = r.FormValue("email")

	rows, err := db.Query("SELECT * FROM userMod")
	checkError.ErrorContr(err, "2")
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
		beeeb.BebMessage("😊", "Registration Successful✅\n ‣ login\n ‣ protected\n")
		defer db.Close()
		bosSignup(w, r)
		db.Close()
	default:
		fmt.Fprintln(w, "record failed error email!! ", uname)
	}

}
