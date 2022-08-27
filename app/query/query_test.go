package query

import "testing"
func TestCheckUser(t *testing.T){
	user := UserValues{
		UserName: "test",
		PassWord: "test1",
	}
	CheckUser(user)
	// Output: [4 6]

}