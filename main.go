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
	http.HandleFunc("/", req_handler.tophandle)
	http.ListenAndServe(":8080", nil)

}
