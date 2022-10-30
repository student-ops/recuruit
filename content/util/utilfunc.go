package util

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net/http"
	"os"
	"strconv"
    "test/query"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Encrypt(input string)(output string) {
    // 暗号化したい文字列
    plaintext := []byte(input)
    // 暗号化された文字列を渡すと、plaintextは渡された文字列になります。
    if len(os.Args) > 1 {
        plaintext = []byte(os.Args[1])
    }

    // aesの暗号化文字列
    key_text := "astaxie12798akljzmknm.ahkjkljl;k"
    if len(os.Args) > 2 {
        key_text = os.Args[2]
    }


    // 暗号化アルゴリズムaesを作成
    c, err := aes.NewCipher([]byte(key_text))
    if err != nil {
        fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
        os.Exit(-1)
    }

    // 暗号化文字列
    cfb := cipher.NewCFBEncrypter(c, commonIV)
    ciphertext := make([]byte, len(plaintext))
    cfb.XORKeyStream(ciphertext, plaintext)

    // 復号文字列
    output = string(ciphertext)
    return output
}
func CheckCookie(w http.ResponseWriter,r *http.Request) string{
    c, err := r.Cookie("user")
    if err != nil{
        return "";
    }
    fmt.Println(c.Value)
    userid := c.Value
    return userid
}
func CheckCookieInt(w http.ResponseWriter,r *http.Request) int{
    cookie_value := CheckCookie(w,r)
    cookie_value_int,_ := strconv.Atoi(cookie_value)
    return cookie_value_int
}

var language map[int]string
func ConvertLang(langd int)(lang string){
    if language == nil{
        language = map[int]string{}
    }
    language[1] = "golang"
    language[2] = "javascript"
    lang = language[langd]
    return lang
}

type ViewMessage struct{
    NameFrom string
    Time string
    Content string
}

func MessageToViewMessage(input query.Message)ViewMessage{
    var result ViewMessage
    userid :=int64(input.IdFrom)
    username := query.CheckUser(userid)
    result.NameFrom = username
    result.Time = input.Time.Format("2006-01-02 15:04:05")
    result.Content = input.Content
    return result
}