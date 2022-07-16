package req_handle

import (
	"net/http"
	"test/query"
	"test/tpl_handle"
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
	if tpl_handle.UserValues == nil{
	tpl_handle.UserValues = make(map[string]string)
	}
	tpl_handle.UserValues["user_id"] = r.FormValue("user_id")
	tpl_handle.UserValues["pass_word"] = r.FormValue("pass_word")
	query.Init()
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