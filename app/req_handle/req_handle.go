package req_handle

import (
	"fmt"
	"net/http"
	"test/query"
	"test/util"
	"text/template"
)

/*
var FormValue map[string]string = map[string]string{
	"user_id": "",
	"pass_word": "",
}
*/
func Top(w http.ResponseWriter, r *http.Request) {
	if userid := util.CheckCookie(w,r); userid != ""{
		threads,_ := query.CheckAllThreads();
		type top_after_value struct{
			Userid string
			Threads  []query.ThreadsVuewer
		}
		Value := top_after_value{
			Userid: userid,
			Threads: threads,
		}
		t, _ := template.ParseFiles(
		"html/top_after.gohtml",
		"html/header.gohtml",
		"html/threads.gohtml",
		"html/footer.gohtml",
		)
		err := t.ExecuteTemplate(w,"top_after.gohtml",Value);if err != nil{
			fmt.Println("error occured in top after")
			panic(err)
		}
	}else{
		t, err := template.ParseFiles("html/top.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(w, nil); err != nil {
				panic(err.Error())
			}
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
func UserConfirm(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/create_account_confirm.html")
	values := map[string]string{
		"user_id": r.FormValue("user_id"),
		"pass_word": r.FormValue("pass_word"),
		"hid_userid": r.FormValue("user_id"),
		"hid_password": r.FormValue("pass_word"),
	}
	fmt.Println(values)
	t.ExecuteTemplate(w,"create_account_confirm.html",values)
}

//create htmlから受け取ってquery.registerに
func UserRegister(h http.HandlerFunc)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request,){
		uservalues := query.UserValues{}
		uservalues.UserName = r.FormValue("hid_userid")
		uservalues.PassWord= r.FormValue("hid_password")
		query.Register(uservalues)
		h(w,r)
	}
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
	uservalue := query.UserValues{
		UserName : r.FormValue("userid"),
		PassWord : r.FormValue(("password")),
	}
	ans := query.UserValues{}
	ans ,err := query.CheckUser(uservalue)
	if err != nil{
		fmt.Printf("ログインに失敗しました。")
	}
	if ans.Userid  != ""{
		authentication := http.Cookie{
			Name: "user_authentication",
			Value: ans.Userid,
			HttpOnly: true,
		}
		http.SetCookie(w,&authentication)
		http.Redirect(w,r,"/",302)
	}else{
		t, _ := template.ParseFiles("html/error.html")
		error_massage := "UserIdまたはパスワードが異なります。"
		t.ExecuteTemplate(w,"error.html",error_massage)
	}
}
func CreateProject(w http.ResponseWriter,R *http.Request){
	t, err := template.ParseFiles("html/create_project.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
func ConfirmProject(h http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		thread := query.Threads{}
		thread.Userid =	util.CheckCookie(w,r);
		thread.Title = r.FormValue("title")
		thread.Detail = r.FormValue("datail")
		thread.Lang = r.FormValue("lang")
		fmt.Println(thread)
		query.ThreadAdd(thread)
		h(w,r)
	}
}
func ThreadPage(w http.ResponseWriter,r *http.Request){
	thread_userid := r.FormValue("hid_userid")
	thread_title := r.FormValue("hid_title")
	thread_date_created := r.FormValue("hid_date_created")
	type thread_value struct{
		Thread query.Threads
		HidUserid string
		HidTitle string	
		HidDateCreated string
	}
	q,_ := query.CheckThread(thread_userid,thread_title,thread_date_created);
	//
	thread_page_value := thread_value{
		Thread: q,
		HidUserid: q.Userid,
		HidTitle: q.Title,
		HidDateCreated: q.Datecreated,
	}
	fmt.Println(thread_page_value)
	t,err := template.ParseFiles("html/thread.gohtml","html/header.gohtml","html/thread_content.gohtml","html/footer.gohtml")
	if err != nil {
		panic(err.Error())
	}
	if err := t.ExecuteTemplate(w, "thread.gohtml",thread_page_value); err != nil {
		panic(err.Error())
	}
}

