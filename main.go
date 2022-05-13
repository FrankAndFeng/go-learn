package main

import (
	"fmt"
	"go-learn/cmd"
)
func main()  {
	fmt.Println("before cmd execute")
	cmd.Execute()
	fmt.Println("after cmd execute")
}