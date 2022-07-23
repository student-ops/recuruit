package query

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
type UserValues struct{
	Userid string
	Password string
	Created string
}
type Threads struct{
	Title string
	Userid string
    Datecreated string
    Lang string
    Detail string
}
var Db *sql.DB
func DbConection(){
	var err error
	Db, err = sql.Open("postgres", "user=recuruit dbname=recuruit password=recuruit sslmode=disable")
	if err != nil {
			panic(err)
	}
	fmt.Println("connected")
}
func (user *UserValues)Register() {
	var err error
	DbConection()
	sql_statement := "INSERT INTO UserValues(userid,password,created)values($1,$2,now());"
	_, err = Db.Exec(sql_statement,user.Userid,user.Password);
	if err != nil{
		panic(err)
	}
}

func CheckUser(userid string)(user UserValues,err error){
	user = UserValues{}
	DbConection()
	sql_statement:= "SELECT userid, password, created FROM UserValues WHERE userid = $1"
	err = Db.QueryRow(sql_statement,userid).Scan(&user.Userid,&user.Password,&user.Created)
	return
}
func CheckThreads()(threads Threads,err error){
	threads = Threads{}
	DbConection()
	sql_statement:="SElECT * from threads"
	err = Db.Query(sql_statement).Scan(&threads.Title,&threads.Userid,&threads.Datecreated,&threads.Lang,&threads.Detail);
	fmt.Println(threads)
	return

}