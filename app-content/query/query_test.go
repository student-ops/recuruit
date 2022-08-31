package query

import (
	"testing"
)
func Test(t *testing.T){
	// user := UserValues{
	// 	UserName: "test",
	// 	PassWord: "test1",
	// }
	// username := CheckUser(3)
	// fmt.Println(username)
	thread := Threads{
		UserId: 1,
		Title: "hell",
		Lang: 2,
		Detail: "hello",
	}
	ThreadAdd(thread)
}