package req_handle

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"test/query"
	"test/session"
	"test/util"
)

type vue_user_profile struct{
	UserName string
	ProfileText string
}

func MyPage(w http.ResponseWriter, r *http.Request,){
	t, err := template.ParseFiles(
		"html/mypage.gohtml",
		"html/header.gohtml",
		"html/footer.gohtml",
	)
	if err != nil{
		log.Fatal(err)
	}
	session_userid := session.Tmp_session[util.CheckCookieInt(w,r)]
	userprofile := vue_user_profile{
		UserName: query.CheckUser(int64(session_userid)),
		ProfileText: query.SelectProfile(int64(session_userid)),
	}
	t.ExecuteTemplate(w,"mypage.gohtml",userprofile)
}
func ChangeProfile(w http.ResponseWriter,r *http.Request){
	text := r.FormValue("user_profile")
	fmt.Printf("req_handle changeprofile text:%s",text)
	userid := session.Tmp_session[util.CheckCookieInt(w,r)]
	query.InsertProfile(userid,text)
	http.Redirect(w,r,"/mypage",302)
}