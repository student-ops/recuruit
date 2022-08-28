package query

import (
	"fmt"

	_ "github.com/lib/pq"
)

type Comments struct{
	ThreadId int64
	Added string
	Userid int64
	CommentContent string
}

func InsertComment(c Comments){
	sql_statement := "INSERT INTO comment (thread_id,added,userid,comment_content VALUES ($1,now(),$2,$3);"
	r , err := Db.Exec(sql_statement,c.ThreadId,c.Userid,c.CommentContent)
	fmt.Println(r)
	if err != nil{
		panic(err)
	}
}

func SelectAllComment(threadid string)(comments []Comments){
	fmt.Println("arrival point0")
	sql_statement := "SELECT * FROM comment WHERE threadid = $1ORDER by added ASC"
	rows,err := Db.Query(sql_statement,threadid); 
	if err != nil{
		fmt.Println("error point 1")
		panic(err)
	}
	for rows.Next(){
		c := Comments{}
		if err = rows.Scan(&c.ThreadId,&c.Added,&c.Userid,&c.CommentContent); err != nil{
			fmt.Println("error point 2")
			panic(err)
		}
		comments = append(comments, c)
	}
	return
}