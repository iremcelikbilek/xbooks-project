package main

import (
	"fmt"
	"net/http"
	"os"

	Login "../Modules/Auth/Login"
	Signup "../Modules/Auth/SignUp"
	AddBook "../Modules/Books"
	fb "../Modules/Firebase"
	UpdatePerson "../Modules/UpdateProfile"
	UserData "../Modules/UserData"
	Util "../Modules/Utils"
)

func main() {
	go fb.ConnectFirebase()
	createServer()
}

func createServer() {
	go http.HandleFunc("/signup", Signup.HandleSignUp)
	go http.HandleFunc("/login", Login.HandleLogin)
	go http.HandleFunc("/books", AddBook.BookAddHandler)
	go http.HandleFunc("/book", AddBook.GetBooks)
	go http.HandleFunc("/updatePerson", UpdatePerson.UpdateProfileHandler)
	go http.HandleFunc("/userData", UserData.UserDataHandler)
	go http.HandleFunc("/userDataSave", UserData.UserDataSaveHandler)
	go http.HandleFunc("/", handleHome)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		print(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	Util.EnableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, mail := Util.CheckToken(r)
	fmt.Fprintf(w, mail)
}
