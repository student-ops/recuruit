package query

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
type UserValues struct{
	UserId int64
	UserName string
	PassWord string
}
type Threads struct{
	ThreadId int64
	Title string
	UserId int64
    DateCreated string
    Lang int
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
	fmt.Println("connected")
}

func Register(user UserValues){
	//user.Password = data.Encrypt
	var err error
	DbConection()
	fmt.Printf("query.register uservalues:%v\n",user)
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
			fmt.Println("query loginckechk succusesed")
		}
	}
	return
}
func CheckUser(userid int64)string {
	DbConection()
	var username string
	sql_statement := "SELECT username from uservalues WHERE userid = $1;"
	if err := Db.QueryRow(sql_statement,userid).Scan(&username); err != nil{
		fmt.Println("err occured checkuser")
		panic(err)
	}
	return username
}
func CheckAllThreads()(threads []Threads,err error){
	DbConection()
	sql_statement := "SELECT * from threads ORDER BY datecreated DESC"
	rows,err :=Db.Query(sql_statement)
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		th :=Threads{}
		if err = rows.Scan(&th.ThreadId, &th.Title,&th.UserId,&th.DateCreated,&th.Lang,&th.Detail);err !=nil{
			// &th.Title,&th.UserId,&th.DateCreated,&th.Lang,&th.Detail
			panic(err)
		}
		threads = append(threads, th)
		fmt.Println(th)
	}
	return
}

func CheckThread(threadid string) (thread Threads,err error){
	DbConection()
	thread = Threads{}
	sql_statement := "SELECT * from threads WHERE threadid = $1"
	err = Db.QueryRow(sql_statement,threadid).Scan(&thread.ThreadId,&thread.Title,&thread.UserId,&thread.DateCreated,&thread.Lang,&thread.Detail)
	if err != nil{
		fmt.Println("can't fetch dbdata in checkthread query")
	}
	return
}
func ThreadAdd(thread Threads){
	DbConection()
	fmt.Printf("thread add thread :%v\n",thread)
	sql_statement := "INSERT INTO threads(threadid,title,userid,datecreated,lang,detail)values(DEFAULT,$1,$2,now(),$3,$4);"
	q, err := Db.Exec(sql_statement, thread.Title, thread.UserId, thread.Lang, thread.Detail);
	if err != nil {
		fmt.Println(q)
		panic(err)
	}
}

func SelectUserIdFromThreadId(threadid string)(userid int64){
	DbConection()
	// if threadid == nil top return nil >>query redirect
	fmt.Printf("query  selectuseridformthreadsid threadid :%v\n",threadid)
	sql_statement := "SELECT userid FROM threads WHERE threadid = $1;"
	err := Db.QueryRow(sql_statement,threadid).Scan(&userid)
	if err != nil{
		log.Fatal(err)
	}
	return
}