package main

import (
	"net/http"
	"text/template"
)
func b(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("top.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
