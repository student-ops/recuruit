package main

import (
	"test/query"
)
type Thread struct{
	name string
	age int
	mon int
}
func main() {
	query.CheckThreads()
}

