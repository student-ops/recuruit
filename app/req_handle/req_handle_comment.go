package req_handle

import (
	"fmt"
	"net/http"
	"strconv"
	"test/query"
	"test/session"
	"test/util"
)

func CommentVuetoComment(c map[string]string)query.Comments{
	comment := query.Comments{}
	comment.ThreadId ,_ = strconv.ParseInt(c["ThreadId"],10,64)
	comment.Userid,_ = strconv.ParseInt(c["UserId"],10,64)
	comment.CommentContent = c["CommentContent"]
	return comment
}
func CommentTOVueComment(c query.Comments)(map[string]string){
	comment := map[string]string{}
	comment["UserName"] = query.CheckUser(c.Userid)
	comment["Added"] = c.Added
	comment["CommentContent"] = c.CommentContent
	return comment
}
func AddComment(w http.ResponseWriter,r *http.Request){
	cookie_user_id :=util.CheckCookie(w,r)
	comment := map[string]string{}
	comment["ThreadId"] = r.FormValue("hid_thread_id")
	comment["UserId"] = cookie_user_id
	comment["CommentContent"] = r.FormValue("comment")
	Comment := CommentVuetoComment(comment)
	fmt.Printf("add Comment %v\n",Comment)
	query.InsertComment(Comment)
	if session.Tmp_thread_session == nil{
		session.Tmp_thread_session = make(map[int]int)	
	}
	cookie_user_id_int,_ := strconv.Atoi(cookie_user_id)
	tmp_thread_id ,_:= strconv.Atoi(r.FormValue("hid_thread_id"))
	session.Tmp_thread_session[cookie_user_id_int] = tmp_thread_id
	http.Redirect(w,r,"/thread_page",302)
}
