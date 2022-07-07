package main
 
import (
    "net/http"
 
    "~/recuruit/req_handler" // 実装したHTTPリクエストハンドラパッケージの読み込み
)
 
func main() {
 
    // "user-form"へのリクエストを関数で処理する
    http.HandleFunc("/user-form", req_handler.HandlerUserForm)
 
    // "user-confirm"へのリクエストを関数で処理する
    http.HandleFunc("/user-confirm", req_handler.HandlerUserConfirm)
 
    // "user-registered"へのリクエストを関数で処理する
    //http.HandleFunc("/user-registered", req_handler.HandlerUserRegistered)
     
    // サーバーを起動
    http.ListenAndServe(":8080", nil)
}
