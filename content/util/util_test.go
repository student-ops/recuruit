package util

import (
	"test/query"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	input := query.Message{
		IdFrom:1,
		IdTo:2,
		Content :"hello",
		Time : time.Now(),
	}

	fmt.Printf("input:%v\n",input)
	fmt.Printf("message:%v\n",MessageToViewMessage(input))
}