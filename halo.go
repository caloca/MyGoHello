package main

import "fmt"

func main(){
  var name string
  
	fmt.Print("What is your name? > ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello " + name +" this is GO World!!\n")
}
