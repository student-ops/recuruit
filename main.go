package main

import (
	/*
		"fmt"
		"test/conf"

		"database/sql"
		_ "github.com/go-sql-driver/mysql"
	*/
	"html/template"
	"net/http"
)
func topHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/top.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",topHandler)



	// サーバーの定義と呼び出し
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}
