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
	UserId string
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
	if _, err = Db.Exec(sql_statement,user.UserName);err != nil{
		panic(err)
	}
	sql_statement = "INSERT INTO personal(userid,password) values (DEFAULT,$1);"
	if _ ,err = Db.Exec(sql_statement,user.PassWord);err != nil{
		panic(err)
	}
}

func LoginCheck(uservalue UserValues)(checked_value UserValues ,err error){
	DbConection()
	checked_value = UserValues{}
	hit_user_id := []int{}
	sql_statement:= "SELECT userid FROM uservalues WHERE username = $1"
	rows ,err := Db.Query(sql_statement,uservalue.UserName)	
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		var tmp int
		rows.Scan(&tmp)
		hit_user_id = append(hit_user_id,tmp)
	}
	fmt.Println(hit_user_id)
	sql_statement = "SELECT password FROM personal WHERE userid = $1"
	for s,_ := range hit_user_id{
		var tmp string
		err = Db.QueryRow(sql_statement,hit_user_id[s]).Scan(&tmp)
		if err!= nil{
			panic(err)
		}
		if tmp == uservalue.PassWord{
			checked_value.UserId = int64(hit_user_id[s])
			checked_value.PassWord = tmp
			fmt.Println(checked_value)
		}
	}
	return 
}
func CheckUser(userid int64)string {
	var username string
	sql_statement := "SELECT (userid, username) from uservalues WHERE userid = $1"
	if err := Db.QueryRow(sql_statement,userid).Scan(&username); err != nil{
		panic(err)
	}
	return username
}
func CheckAllThreads()(threads []Threads,err error){
	DbConection()
	rows,err :=Db.Query("SElECT * from threads ORDER BY datecreated DESC")
	if err !=nil{return}
	for rows.Next(){
		th :=Threads{}
		if err = rows.Scan(&th.Title,&th.UserId,&th.Datecreated,&th.Lang,&th.Detail);err !=nil{return}
		threads = append(threads, th)
	}
	return
}
func CheckThread(userid string,title string,datecreated string) (thread Threads,err error){
	fmt.Println("arrival check thred")
	DbConection()
	thread = Threads{}
	rows := "SELECT title,userid,datecreated,lang,detail from threads WHERE userid = $1"
	err = Db.QueryRow(rows,userid).Scan(&thread.Title,&thread.UserId,&thread.Datecreated,&thread.Lang,&thread.Detail)
	if err != nil{
		fmt.Println("can't fetch dbdata in checkthread query")
	}
	return
}
func ThreadAdd(thread Threads)error {
	var err error
	sql_statement := "INSERT INTO threads(title,userid,datecreated,lang,detail)values($1,$2,now(),$3,$4);"
	_ , err = Db.Exec(sql_statement, thread.Title,thread.UserId,thread.Lang,thread.Detail);
	if err != nil {
		panic(err)
	}
	return err
}