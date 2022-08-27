package query

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
type UserValues struct{
	UserId int64
	UserName string
	PassWord string
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
	Db, err = sql.Open("postgres", "port =5432 user=recuruit dbname=recuruit password=recuruit sslmode=disable")
	if err != nil {
		fmt.Println("can't connect sql")
			panic(err)
	}
}

func Register(user UserValues){
	//user.Password = data.Encrypt
	var err error
	DbConection()
	sql_statement := "INSERT INTO uservalues(userid,username,created)values(DEFAULT,$1,now());"
	_, err = Db.Exec(sql_statement,user.UserName);
	sql_statement = "INSERT INTO personal(userid,password) values (DEFAULT,$1);"
	_ ,err = Db.Exec(sql_statement,user.PassWord);
	if err != nil{
		panic(err)
	}
}

func CheckUser(userid string)(checked_value UserValues){
	DbConection()
	checked_value = UserValues{}
	sql_statement:= "SELECT userid, username FROM uservalues WHERE userid = $1"
	err := Db.QueryRow(sql_statement,userid).Scan(&checked_value.UserId,&checked_value.UserName)
	if err != nil {
		panic(err)
	}
	sql_statement := "SLECT password FROM "

	return
}

func CheckAllThreads()(threads []Threads,err error){
	DbConection()
	rows,err :=Db.Query("SElECT * from threads ORDER BY datecreated DESC")
	if err !=nil{return}
	for rows.Next(){
		th :=Threads{}
		if err = rows.Scan(&th.Title,&th.Userid,&th.Datecreated,&th.Lang,&th.Detail);err !=nil{return}
		threads = append(threads, th)
		fmt.Println(th)
	}
	return
}
func CheckThread(userid string,title string,datecreated string) (thread Threads,err error){
	fmt.Println("arrival check thred")
	DbConection()
	thread = Threads{}
	rows := "SELECT title,userid,datecreated,lang,detail from threads WHERE userid = $1"
	err = Db.QueryRow(rows,userid).Scan(&thread.Title,&thread.Userid,&thread.Datecreated,&thread.Lang,&thread.Detail)
	if err != nil{
		fmt.Println("can't fetch dbdata in checkthread query")
	}
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