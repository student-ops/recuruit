package query

import (
	"database/sql"
	"fmt"
	"test/util"

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
	//user.Password = data.Encrypt
	var err error
	DbConection()
	encpath := util.Encrypt(user.Password)
	sql_statement := "INSERT INTO UserValues(userid,password,created)values($1,$2,now());"
	_, err = Db.Exec(sql_statement,user.Userid,encpath);
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

func CheckThreads()(threads []Threads,err error){
	DbConection()
	rows,err :=Db.Query("SElECT * from threads")
	if err !=nil{return}
	for rows.Next(){
		th :=Threads{}
		if err = rows.Scan(&th.Title,&th.Userid,&th.Datecreated,&th.Lang,&th.Detail);err !=nil{return}
		threads = append(threads, th)
	}
	fmt.Println(threads)
	return
}

func ThreadAdd(thread Threads){
	thread.Userid = "hurry"
	fmt.Println(thread)
	var err error
	sql_statement := "INSERT INTO threads(title,userid,datecreated,lang,detail)values($1,$2,now(),$3,$4);"
	_ , err = Db.Exec(sql_statement, thread.Title,thread.Userid,thread.Lang,thread.Detail);
	if err != nil {
		panic(err)
	}
	return
}