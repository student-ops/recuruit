package req_handler

import (
	"net/http"
	"text/template"
)
func topHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("top.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
