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
	mux.HandleFunc("/user_confirm",req_handle.UserConfirm)
	mux.HandleFunc("/user_register",req_handle.UserRegister(req_handle.Top))
	mux.HandleFunc("/login_confirm",req_handle.LoginConfirm)
	mux.HandleFunc("/create_project",req_handle.CreateProject)
	mux.HandleFunc("/confirm_project",req_handle.ConfirmProject(req_handle.LoginConfirm))
	mux.HandleFunc("/vue_thread",req_handle.VueThread)
	// サーバーの定義と呼び出し
	server := &http.Server{
		Addr: "localhost:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}
