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
	ans := query.UserValues{}
	ans ,_ = query.CheckUser(uservlues.Userid)
	fmt.Println(ans)
	if ans.Userid  != ""{

		t, _ := template.ParseFiles("html/top_after.html","html/post.html")
		t.Execute(w, ans.Userid)
	}
}
func Post(w http.ResponseWriter,R *http.Request){
	threds := query.
}