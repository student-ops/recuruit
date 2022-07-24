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
	Userid := r.FormValue("userid")
	ans := query.UserValues{}
	ans ,_ = query.CheckUser(Userid)
	fmt.Println(ans)
	threads := []query.Threads{}
	threads,_ = query.CheckThreads();
	if ans.Userid  != ""{
		t, _ := template.ParseFiles("html/top_after.html","html/post.html")
		err := t.ExecuteTemplate(w,"top_after",ans.Userid);if err != nil{panic(err)}
		t.ExecuteTemplate(w,"post", threads)
	}
}
func Posts(w http.ResponseWriter, r *http.Request){
	threads := []query.Threads{}
	threads,_ = query.CheckThreads();
	t, err := template.ParseFiles("html/post.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, threads,); err != nil {
		panic(err.Error())
	}
}