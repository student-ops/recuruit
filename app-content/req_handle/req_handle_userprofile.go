package req_handle

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"test/query"
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
	userid := util.CheckCookie(w,r)
	d ,_:=  strconv.ParseInt(userid,10,64)
	userprofile := vue_user_profile{
		UserName: query.CheckUser(d),
		ProfileText: query.SelectProfile(d),
	}
	fmt.Printf("mypage userid :%s",userid)
	t.ExecuteTemplate(w,"mypage.gohtml",userprofile)
}

func ChangeProfile(){

}