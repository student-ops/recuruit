package query

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
var Db *sql.DB
func Init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
			panic(err)
	}
	fmt.Println("接続成功")
}