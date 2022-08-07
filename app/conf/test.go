package main
import(
	"test/query"
)
func main(){
	user := query.UserValues{"hurry","kane","now()"}
	user.Register()
}