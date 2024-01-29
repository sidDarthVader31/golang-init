package main

import (
	"fmt"
	"time"
)

func main(){
  go printRoutines(5)
  printNumbers(5)
}

func printNumbers(a int){
  for i := 0 ;i<a;i ++{
    fmt.Println(i)
  }
}

func printRoutines(a int) {
  for i :=0 ; i< a; i++{
    time.Sleep(100 * time.Millisecond)
    fmt.Println("i::", i)
  }
}
