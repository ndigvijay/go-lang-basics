package main

import "fmt"


func SayGreeting(name string){
	fmt.Printf("yo hello there %v \n",name);
}
func SayBye(name string){
	fmt.Printf("yo good bye have a great day  %v \n",name)
}

func main(){
	SayGreeting("ndv")
	SayGreeting("aavish")
	SayBye("aavish")
	SayBye("ndv")
}