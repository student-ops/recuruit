package req_handle

import (
	"net/http"
	"strconv"
	"test/query"
	"test/util"
)

func CommentVuetoComment(c map[string]string)query.Comments{
	comment := query.Comments{}
	comment.ThreadId ,_ = strconv.ParseInt(c["ThreadId"],10,64)
	comment.Userid,_ = strconv.ParseInt(c["UserId"],10,64)
	comment.CommentContent = c["ContentContent"]
	return comment
}
func CommentTOVueComment(c query.Comments)(map[string]string){
	comment := map[string]string{}
	comment["UserName"] = query.CheckUser(c.Userid)
	comment["Added"] = c.Added
	comment["CommentContet"] = c.CommentContent
	return comment
}
func AddComment(w http.ResponseWriter,r *http.Request){
	comment := map[string]string{}
	comment["ThreadId"] = r.FormValue("hid_thread_id")
	comment["UserId"] = util.CheckCookie(w,r)
	comment["ContentContent"] = r.FormValue("comment")
	Comment := CommentVuetoComment(comment)
	query.InsertComment(Comment)
}

