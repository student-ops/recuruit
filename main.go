package main
import(
	"fmt"
	"test/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func main(){
	confDB, err := conf.ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 設定値から接続文字列を生成
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName)

	// データベース接続
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	// deferで処理終了前に必ず接続をクローズする
	defer db.Close()
	// 接続確認
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	}else {
		fmt.Println("データベース接続成功！")
	}
}


