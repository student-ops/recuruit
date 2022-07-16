package query

import (
	"database/sql"
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
	Db, err = sql.Open("postgres", "user=recuruit dbname=recuruit password=recuruit sslmode=disable")
	if err != nil {
			panic(err)
	}
	fmt.Println("接続成功")
	fmt.Println(UserValues)
	sql_statement := "INSERT INTO test(name)values($1);"
	_, err = Db.Exec(sql_statement,UserValues["user_id"]);
	if err != nil{
		panic(err)
	}
}