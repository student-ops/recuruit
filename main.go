package main

import (
	/*
		"fmt"
		"test/conf"

		"database/sql"
		_ "github.com/go-sql-driver/mysql"
	*/
	"net/http"
	"test/req_handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/static",tophandler.topHandler)

	// サーバーの定義と呼び出し
	server := &http.Server{
		Addr: "0.0.0.0.:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}
