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
type THredsVuewer struct{
	Title string
	Userid string
    Datecreated string
    Lang string
    Detail string
	HidUserid string
	HidTitle string
	HidDateCreated string
}

var Db *sql.DB
func DbConection(){
	var err error
	Db, err = sql.Open("postgres", "host = db port =5432 user=recuruit dbname=recuruit password=recuruit sslmode=disable")
	if err != nil {
		fmt.Println("can't connect sql")
			panic(err)
	}
}

func Register(user UserValues){
	//user.Password = data.Encrypt
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

func CheckAllThreads()(threads []THredsVuewer,err error){
	DbConection()
	rows,err :=Db.Query("SElECT * from threads ORDER BY datecreated DESC")
	if err !=nil{return}
	for rows.Next(){
		th :=THredsVuewer{}
		if err = rows.Scan(&th.Title,&th.Userid,&th.Datecreated,&th.Lang,&th.Detail);err !=nil{return}
		th.HidUserid = th.Userid
		th.HidTitle = th.Title
		th.HidDateCreated = th.Datecreated
		threads = append(threads, th)
		fmt.Println(th)
	}
	return
}
func CheckThread(userid string,title string,datecreated string) (thread Threads,err error){
	fmt.Println("arrival check thred")
	fmt.Println(userid)
	fmt.Println(title)
	fmt.Println(datecreated)
	DbConection()
	thread = Threads{}
	rows := "SELECT title,userid,datecreated,lang,detail from threads WHERE userid = $1"
	err = Db.QueryRow(rows,userid).Scan(&thread.Title,&thread.Userid,&thread.Datecreated,&thread.Lang,&thread.Detail)
	return
}
func ThreadAdd(thread Threads){
	var err error
	sql_statement := "INSERT INTO threads(title,userid,datecreated,lang,detail)values($1,$2,now(),$3,$4);"
	_ , err = Db.Exec(sql_statement, thread.Title,thread.Userid,thread.Lang,thread.Detail);
	if err != nil {
		panic(err)
	}
	return
}