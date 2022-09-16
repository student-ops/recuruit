package query

import (
	"fmt"
	"testing"
)
func Test(t *testing.T){
	// result := SelectNewestMessage(1)
	// edited := DeleteSameRootMessage(result)
	// fmt.Printf("--------------\nfetched \n")
	// for index,i := range result{
	// 	fmt.Printf("index:%d,content:%v\n",index,i)
	// }
	// fmt.Println("--------------")
	// for index,i := range edited{
	// 	fmt.Printf("index:%d,content:%v\n",index,i)
	// }
	result := SelectIndivisualMessage(1,2)
	for _,i := range result{
		fmt.Println(i)
	}
}