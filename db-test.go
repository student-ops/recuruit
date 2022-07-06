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
	db, err := sql.Open("mysql", "lakky1:lakky1@(localhost:3306)/test")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()
	
	ins, err := db.Prepare("INSERT INTO user VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// SQLの実行
	res, err := ins.Exec(2, "chuken")
	if err != nil {
		log.Fatal(err)
	}

	// 結果の取得
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastInsertID)
}
