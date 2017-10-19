package main

import (
  "fmt"
  "github.com/caloca/stringutil"
)

func main(){
  var name string

	fmt.Print("What is your name? > ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello " + name +" this is GO World!!\n")
	fmt.Printf(stringutil.Reverse(name) + "\n")

}
