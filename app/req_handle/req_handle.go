package req_handle

import (
	"fmt"
	"net/http"
	"strconv"
	"test/query"
	"test/util"
	"text/template"
)
type vue_threads struct{
	Title string
	UserName string
	Lang string
	Detail string
	DateCreated string
	HidThreadId string
}
func ThreadToVueThread(thread_data query.Threads)vue_threads{
	threadid  := strconv.FormatInt(thread_data.ThreadId ,10)
	username  := query.CheckUser(thread_data.UserId)
	vue_thread := vue_threads{
		thread_data.Title,
		username,
		thread_data.Lang,
		thread_data.Detail,
		thread_data.DateCreated,
		threadid,
	}
	return vue_thread
}
func Top(w http.ResponseWriter, r *http.Request) {
	if userid := util.CheckCookie(w,r); userid != ""{
		threads,_ := query.CheckAllThreads();
		fmt.Println(threads)
		//you should make this map. this declare is toolong
		Thread := []vue_threads{}
		for i := range threads {		th := ThreadToVueThread(threads[i]);
			Thread = append(Thread,th)
		}
		type top_after_value struct{
			Userid string
			Threads  []vue_threads
		}
		userid_int ,_:= strconv.ParseInt(userid,10,64)	
		username := query.CheckUser(userid_int)
		Value := top_after_value{
			Userid: username,
			Threads: Thread,
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
		UserName : r.FormValue("username"),
		PassWord : r.FormValue("password"),
	}
	fmt.Println(uservalue)
	ans := query.UserValues{}
	ans ,err := query.LoginCheck(uservalue)
	fmt.Println(ans)
	if err != nil{
		fmt.Println("ログインに失敗しました。")
	}
	// decelrede but it peform false
	if  ans.UserId != 0{
		ans.UserName = uservalue.UserName
		authentication := http.Cookie{
			Name: "user_authentication",
			Value:strconv.FormatInt( ans.UserId,10),
			HttpOnly: true,
		}
		http.SetCookie(w,&authentication)
		http.Redirect(w,r,"/",302)
	}else{
		t, _ := template.ParseFiles("html/error.html")
		m := make(map[string]string)
		m["ErrorMassage"]= "UserIdまたはパスワードが異なります。"
		fmt.Println(m["ErrorMassage"])
		t.ExecuteTemplate(w,"error.html",m)
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
		s, _ := strconv.ParseInt(util.CheckCookie(w,r), 10, 64)
		thread := query.Threads{}
		thread.UserId  = s
		thread.Title = r.FormValue("title")
		thread.Detail = r.FormValue("datail")
		thread.Lang = r.FormValue("lang")
		fmt.Println(thread)
		query.ThreadAdd(thread)
		h(w,r)
	}
}
func ThreadPage(w http.ResponseWriter,r *http.Request){
	thread_id := r.FormValue("hid_thread_id")
	th,_ := query.CheckThread(thread_id);
	type thread_page struct{
		Title string
		UserName string
		DateCreated string
		Lang string
		Detail string
		HidThreadId string
		Comments []map[string]string
	}
	thread := ThreadToVueThread(th)
	var Comments  []map[string]string
	comments := query.SelectAllComment(thread_id)
	for i := range comments{
		s := map[string]string{}
		Comments[i] = s
	}
	thread_page_value := thread_page{
		thread.Title,
		thread.UserName,
		thread.DateCreated,
		thread.Lang,
		thread.Detail,
		thread.HidThreadId,
		Comments,
	}
	fmt.Printf("thread_page_value\n %v\n",thread_page_value)
	t,err := template.ParseFiles("html/thread.gohtml","html/header.gohtml","html/comments.gohtml","html/footer.gohtml")
	if err != nil {
		panic(err.Error())
	}
	if err := t.ExecuteTemplate(w, "thread.gohtml",thread_page_value); err != nil {
		panic(err.Error())
	}
}
