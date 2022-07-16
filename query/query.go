package query

import (
	"database/sql"
	"errors"
	"fmt"
	"test/tpl_handle"

	_ "github.com/lib/pq"
)
var Db *sql.DB
func Init() {
	var err error
	var UserValues map[string]string
	if tpl_handle.UserValues == nil{
		tpl_handle.UserValues = make(map[string]string)
	}
	UserValues = tpl_handle.UserValues
	DbConection()
	fmt.Println(UserValues)
	sql_statement := "INSERT INTO test(name)values($1);"
	_, err = Db.Exec(sql_statement,UserValues["user_id"]);
	if err != nil{
		panic(err)
	}
}
func DbConection(){
	var err error
	Db, err = sql.Open("postgres", "user=recuruit dbname=recuruit password=recuruit sslmode=disable")
	if err != nil {
			panic(err)
	}
	fmt.Println("接続成功")
}
func CheckUser(id int,pass string)(bool){
	DbConection()
	sql_statement:= "SELECT id FROM test WHERE id = $1 and name = $2"
	rows, err := Db.Exec(sql_statement,id,pass)
	if err != nil{
		panic(err)
	}
	fmt.Println(rows)
	if rows == nil {
		err := errors.New("UserIdが存在しません。")
		fmt.Println(err)
		return false
	}
	fmt.Println("hit!!")
	return true
}