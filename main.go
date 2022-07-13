package main

import (
	/*
		"fmt"
		"test/conf"

		"database/sql"
		_ "github.com/go-sql-driver/mysql"
	*/
	"net/http"
)
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",b)
	// サーバーの定義と呼び出し
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}