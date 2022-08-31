package req_handle

import(
	"net/http"
	"fmt"
)

func AddComment(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello")

}