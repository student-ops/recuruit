package req_handle

import (
	"fmt"
	"net/http"
	"test/query"
	"text/template"
)

/*
var FormValue map[string]string = map[string]string{
	"user_id": "",
	"pass_word": "",
}
*/
func Top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/top.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/create_account.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
func UserConfirm(w http.ResponseWriter, r *http.Request){
	uservlues := query.UserValues{}
	uservlues.Userid = r.FormValue("user_id")
	uservlues.Password= r.FormValue("pass_word")
	uservlues.Register()
}


func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/login.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
func LoginConfirm(w http.ResponseWriter, r *http.Request){
	uservlues := query.UserValues{}
	uservlues.Userid = r.FormValue("userid")
	fmt.Println(uservlues)
	ans ,_ := query.CheckUser(uservlues.Userid)
	fmt.Println(ans)
	t, _ := template.ParseFiles("html/top_after.html")
	t.Execute(w, "Hello World!")
}