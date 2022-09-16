package query

import(
	// "database/sql"
	// "fmt"
	"log"
	"time"
	_ "github.com/lib/pq"
)

type Message struct{
	IdFrom int
	IdTo int
	Content string
	Time time.Time
}

func DeleteSameRootMessage(input []Message)(result []Message){
	var skip []int
	for i_index,i := range input{
		for j := i_index+1;j< len(input);j++{
			if (i.IdFrom == input[j].IdTo  && i.IdTo == input[j].IdFrom) || (i.IdFrom == input[j].IdFrom && i.IdTo == input[j].IdTo){
				// fmt.Printf("i_index:%d j:%d\n",i_index,j)
				if i.Time.After(input[j].Time){
					skip = append(skip,j)
				}else{
					skip = append(skip,i_index)
				}
			}
		}
		skip_check := true
		for _,k := range skip{
			if i_index == k{
				skip_check = false
			}
		}
		if skip_check {
			result = append(result,i)
		}
	}
	return
}

func SelectNewestMessage(userid int)(result []Message){
	DbConection()
	sql_statement := "SELECT DISTINCT ON(follow,followed,time) follow,followed,message,time FROM direct_message WHERE follow = $1 OR followed =$1"
	rows,err := Db.Query(sql_statement,userid)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		message := Message{}
		if err = rows.Scan(&message.IdFrom,&message.IdTo,&message.Content,&message.Time);err != nil{
			log.Fatal(err)
		}
		result = append(result,message)
	}
	return result
}
//リファクタリング sql_statement いかほとんど繰り返し
func SelectIndivisualMessage(from int,to int)(result []Message){
	DbConection()
	sql_statement := "SELECT follow,followed,message,time FROM direct_message WHERE follow =$1 and followed =$2 or follow =$2 and followed =$1 ORDER BY time asc LIMIT 40"
	rows,err := Db.Query(sql_statement,from,to)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		message := Message{}
		if err = rows.Scan(&message.IdFrom,&message.IdTo,&message.Content,&message.Time);err != nil{
			log.Fatal(err)
		}
		result = append(result,message)
	}
	return result
}