package main

import "fmt";

func main(){
  fmt.Println("hello world");

  var ages[3]int = [3]int {20,30,4};
  fmt.Println(ages);
  fmt.Println(len(ages));
  var name =[4]string{"digvijay","nihal","dmello"};
  fmt.Println(name,len(name));
  //slice

  var scores=[]int{10,20,30};
  scores[2]=80;
  var newscores =append(scores,69);
  fmt.Println(newscores , len(newscores));

}
