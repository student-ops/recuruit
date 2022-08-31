package req_handle

import (
	"fmt"
	"net/http"
	"strconv"
	"test/query"
	"test/session"
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
		util.ConvertLang(thread_data.Lang),
		thread_data.Detail,
		thread_data.DateCreated,
		threadid,
	}
	return vue_thread
}
func Top(w http.ResponseWriter, r *http.Request) {
	cookie_value := util.CheckCookie(w,r)
	cookie_value_int ,_ := strconv.Atoi(cookie_value)
	fmt.Printf("top tmp_cookie_value_int %v\n",cookie_value_int)
	fmt.Printf("top tmp_session[userid]:%v\n",session.Tmp_session[cookie_value_int])
	if  session.Tmp_session[cookie_value_int]{
		threads,_ := query.CheckAllThreads();
		//you should make this map. this declare is toolong
		Thread := []vue_threads{}
		for i := range threads {		
			th := ThreadToVueThread(threads[i]);
			Thread = append(Thread,th)
		}
		type top_after_value struct{
			Userid string
			Threads  []vue_threads
		}
		userid_int ,_:= strconv.ParseInt(cookie_value,10,64)	
		username := query.CheckUser(userid_int)
		Value := top_after_value{
			Userid: username,
			Threads: Thread,
		}
		t, _ := template.ParseFiles(
		"html/top_after.gohtml",
		"html/top_header.gohtml",
		"html/threads.gohtml",
		"html/serch.gohtml",
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
// redirect login comfirm
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
	login_check_ans ,err := query.LoginCheck(uservalue)
	fmt.Printf("req_handle login conf ans:%v\n",login_check_ans)
	if err != nil || login_check_ans.UserId == 0{
		fmt.Println("ログインに失敗しました。")
	}
	//query.Logincheck aceepted
	if  login_check_ans.UserId != 0{
		login_check_ans.UserName = uservalue.UserName
		authentication := http.Cookie{
			Name: "user",
			Value: strconv.FormatInt(login_check_ans.UserId,10),//string value
			HttpOnly: true,
		}
		http.SetCookie(w,&authentication)
		if session.Tmp_session == nil{
			session.Tmp_session = make(map[int]bool)
		}
		cookie_value_int := int(login_check_ans.UserId)
		fmt.Printf("reqhandle login conf cookivalueint :%d\n",cookie_value_int)
		session.Tmp_session[cookie_value_int] = true
		fmt.Printf("reqhandle login conf cookivalueint :%v\n",session.Tmp_session[cookie_value_int])
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
	t, err := template.ParseFiles(
		"html/create_project.gohtml",
		"html/header.gohtml",
		"html/footer.gohtml",
	)
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
func ConfirmProject(w http.ResponseWriter,r *http.Request){
		s, _ := strconv.ParseInt(util.CheckCookie(w,r), 10, 64)
		fmt.Printf("confirmproj cookie value:%v\n",s)
		lang,_ :=  strconv.Atoi(r.FormValue("lang"))
		thread := query.Threads{}
		thread.UserId  = s
		thread.Title = r.FormValue("title")
		thread.Detail = r.FormValue("datail")
		thread.Lang = lang
		fmt.Println(thread)
		query.ThreadAdd(thread)
		http.Redirect(w,r,"/",302)
}
var ThreadIdComment string
func ThreadPage(w http.ResponseWriter,r *http.Request){
	thread_id := r.FormValue("hid_thread_id")
	fmt.Printf("threadpage threadid:%v\n",thread_id)
	cookie_value_int := util.CheckCookieInt(w,r)
	if thread_id == ""{
		thread_id = strconv.Itoa(session.Tmp_thread_session[cookie_value_int])
		fmt.Printf("session thread id %s\n",thread_id)
	}else{
		//redirect top
	}
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
	comments := query.SelectAllComment(thread_id)
	fmt.Printf("len(comment):%v",len(comments))
	Comments := make([]map[string]string,len(comments))
	for i := range comments{
		s := map[string]string{}
		s = CommentTOVueComment(comments[i])
		Comments[i] = s
	}
	fmt.Printf("comments;%v",Comments)
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
