package main

import (
	"fmt"
	"time"
)

func main(){
  go print("siddharth")
  print("bisht")
}

func print(s string){
  time.Sleep(100 * time.Millisecond)
  for i:=1;i<5 ;i++{
    fmt.Println(s)
  }
}
