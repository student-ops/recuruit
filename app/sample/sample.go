package sample

import (
	"fmt"
	"html/template"
	"net/http"
	"test/query"
)


func ThreadViwer (w http.ResponseWriter,r *http.Request){
	Q := query.Threads{
		Title: "test html template",
		Userid: "rakky",
		Datecreated: "tody",
		Lang: "japanese",
		Detail: "hello world",
	}
	type thread_page struct{
		Thread query.Threads
		HidThreadId string
	}
	thread_page_value := thread_page{
		Thread: Q,
		HidThreadId: "test",
	}
	fmt.Println(thread_page_value)
	t,err := template.ParseFiles("html/thread.html","html/test_thread.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.ExecuteTemplate(w,"thread.html",thread_page_value); err != nil {
		panic(err.Error())
	}




}