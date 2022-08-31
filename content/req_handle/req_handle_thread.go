package req_handle

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"test/query"
)
type VueUserProfiles struct{
	UserName string
	Profile string
}
func UserPage(w http.ResponseWriter, r *http.Request){
	threadid := r.FormValue("hid_thread_id")
	vueuserprofile := VueUserProfiles{}
	userid := query.SelectUserIdFromThreadId(threadid)
	vueuserprofile.Profile = query.SelectProfile(int64(userid))
	vueuserprofile.UserName = query.CheckUser(userid)
	t,err := template.ParseFiles(
		"html/user_page.gohtml",
		"html/header.gohtml",
		"html/footer.gohtml",
	)
	if err != nil{
		fmt.Printf("req UserPage parse fiel error")
		log.Fatal(err)
	}
	t.ExecuteTemplate(w,"user_page.gohtml",vueuserprofile)
}

