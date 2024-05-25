package main

import (
   "fmt"
)


func main(){
    names:=[]string{"dmells","ndv","nyhal"};
    for index,value := range names{
	fmt.Printf("index and value are %v is %v\n",index,value);
	}

}
