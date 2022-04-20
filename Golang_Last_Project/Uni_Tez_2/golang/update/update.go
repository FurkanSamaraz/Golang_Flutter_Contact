package update

import (
	"encoding/json"
	"fmt"
	"net/http"

	helper "github.com/FurkanSamaraz/IsEmpty"
	"main.go/database"
)

type UserModel struct {
	Id       int
	Username string
	Password string
	Email    string
	Token    string
}

type Names struct {
	Successful string
	Message    string
	UserModel  UserModel
}

var uname, pwd, email string

func bosSignup(w http.ResponseWriter, r *http.Request) {

	unameCheck := helper.IsEmpty(uname)
	pwdCheck := helper.IsEmpty(pwd)
	mailCheck := helper.IsEmpty(email)

	if unameCheck || pwdCheck || mailCheck {

	} else {
		fmt.Fprintf(w, "Error Empty! \n")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	var userr Names
	db := database.OpenConnention()
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
