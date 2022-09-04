package query

import (
	"fmt"

	_ "github.com/lib/pq"
)

func SelectThreadFromUser(userid int)[]Threads{
	DbConection()
	sql_statement := "SELECT * FROM threads WHERE userid = $1 ORDER BY datecreated DESC"
	rows,err := Db.Query(sql_statement,userid)
	threads := []Threads{}
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		th :=Threads{}
		if err = rows.Scan(&th.ThreadId, &th.Title,&th.UserId,&th.DateCreated,&th.Lang,&th.Detail);err !=nil{
			panic(err)
		}
		threads = append(threads, th)
		fmt.Println(th)
	}
	fmt.Printf("selectthreadsfrom user []Threads:%v",threads)
	return threads
}