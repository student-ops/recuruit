package req_handle

import(
	// "http/template"
	"net/http"
	"test/util"
	"test/query"
	"fmt"
)

func Message(w http.ResponseWriter,r *http.Request){
	fmt.Println("arrival message")
	cookie_user_id := util.CheckCookieInt(w,r)
	fmt.Printf("message userid %d",cookie_user_id)
	massge_list := query.DeleteSameRootMessage(query.SelectNewestMessage(cookie_user_id))
	value util.ViewMessage
}

