package main

import (
	"fmt"
	"test/query"
)
func main(){
	user := query.UserValues{Userid: "mike",Password:"hello"}
	user.Register()
	ans ,_ := query.CheckUser("ryuta")
	fmt.Println(ans)
}
