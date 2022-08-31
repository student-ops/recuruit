package query

import "log"
type UserProfile struct{
	UserId string
	ProFileText string
}


func InsertProfile(user_id int, text string){
	DbConection()
	sql_statement := "INSERT INTO userprofile(userid,profiletext)values ($1,$2);"
	_,err := Db.Exec(sql_statement,user_id,text)
	if err != nil{
		log.Fatal(err)
	}
}

func SelectProfile(user_id int64)(text string){
	DbConection()
	sql_statement := "SELECT profiletext FROM userprofile WHERE userid = $1"
	rows, err := Db.Query(sql_statement,user_id)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		if err =rows.Scan(&text); nil != err{
			log.Fatal(err)
		}
	}
	return text
}

