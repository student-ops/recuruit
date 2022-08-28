package query

import (
	"fmt"
	"testing"
)
func Test(t *testing.T){
	// user := UserValues{
	// 	UserName: "test",
	// 	PassWord: "test1",
	// }
	// username := CheckUser(3)
	// fmt.Println(username)
	th ,_ := CheckAllThreads();
	fmt.Println(th)
}