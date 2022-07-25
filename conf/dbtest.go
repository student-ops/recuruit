package main

import (
	/*
		"fmt"
		"test/conf"
		"database/sql"
		_ "github.com/go-sql-driver/mysql"
	*/
	"fmt"
	"net/http"
)
func Hello(w http.ResponseWriter,r *http.Request){
	fmt.Println("w:")
	fmt.Println(w)
	fmt.Println("r:")
	fmt.Println(r)
}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello",Hello)
		server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}