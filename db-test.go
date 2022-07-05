package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectOnly()
}

func connectOnly() {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "lakky1:lakky1@(localhost:3306)/db1")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了")
	}
}
