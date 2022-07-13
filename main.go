package main

import (
	/*
		"fmt"
		"test/conf"

		"database/sql"
		_ "github.com/go-sql-driver/mysql"
	*/
	"net/http"
	"test/req_handle"
)
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",req_handle.Top)
	mux.HandleFunc("/create_account",req_handle.CreateAccount)
	mux.HandleFunc("/login",req_handle.Login)
	// サーバーの定義と呼び出し
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}