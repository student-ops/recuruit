package main
 
import (
    "database/sql"
    "fmt"
 
    "test/conf"  // 実装した設定パッケージの読み込み
    "test/query" // 実装したクエリパッケージの読み込み
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
 
    // 設定ファイルを読み込む
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
 
    // INSERTの実行
    id, err := query.InsertUser("USER-0001", "山田 hoge郎", "pass", db)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("登録されたユーザのIDは【%d】です。\n", id)
 
    // SELECTの実行
    user, err := query.SelectUserById(id, db)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("SELECTされたユーザ情報は以下の通りです。\n")
    fmt.Printf("[ID] %s\n", user.Id)
    fmt.Printf("[アカウント] %s\n", user.Account)
    fmt.Printf("[名前] %s\n", user.Name)
    fmt.Printf("[パスワード] %s\n", user.Passwd)
    fmt.Printf("[登録日] %s\n", user.Created)
} 
