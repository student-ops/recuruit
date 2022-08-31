package query

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Comments struct{
	ThreadId int64
	Added string
	Userid int64
	CommentContent string
}

func InsertComment(c Comments){
	DbConection()
	sql_statement := "INSERT INTO comment (threadid,added,userid,comment_content) VALUES ($1,now(),$2,$3);"
	r , err := Db.Exec(sql_statement,c.ThreadId,c.Userid,c.CommentContent)
	if err != nil{
		fmt.Printf("db exec result %v\n",r)
		panic(err)
	}
}

func SelectAllComment(threadid string)(comments []Comments){
	sql_statement := "SELECT * FROM comment WHERE threadid = $1"
	fmt.Printf("threadid for comment %s \n",threadid)
	rows,err := Db.Query(sql_statement,threadid)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("rows:%v\n",rows)
	defer rows.Close()
	for rows.Next(){
		c := Comments{}
		err = rows.Scan(&c.ThreadId,&c.Added,&c.Userid,&c.CommentContent)
		if err != nil{
			panic(err)
		}
		comments = append(comments, c)
	}
		fmt.Println("not yet")
	return
}