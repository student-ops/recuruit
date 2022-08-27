package req_handle

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
    var intPointer *int

    // To set the value use:
    // intValue := 3
    // intPointer = &intValue

    if intPointer == nil {
        fmt.Println("The variable is nil")
    } else {
        fmt.Printf("The variable is set to %v\n", *intPointer)
    }
}