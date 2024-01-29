package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main(){
  fmt.Println("main exection starts")
  fmt.Println("number of CPus:", runtime.NumCPU())
  fmt.Println("number of go routines:", runtime.NumGoroutine())
  fmt.Println("OS:", runtime.GOOS)
  fmt.Println("Arch:", runtime.GOARCH)
  // go print("siddharth")
  // fmt.Println("no of go routines now:", runtime.NumGoroutine())
  // print("bisht")

  var wg sync.WaitGroup
  wg.Add(1)
  go f1(&wg)
  f2()

  wg.Wait()

  fmt.Println("main execution ends")
}


func f1(wg *sync.WaitGroup){
  for i:=1;i<5;i++{
    time.Sleep(time.Second)
    fmt.Println("f1()i:", i)
  }
  (*wg).Done()
}
func f2(){
  for i:=6;i<=10;i++{
    fmt.Println("f2()i:", i)
  }
}


func print(s string){
  time.Sleep(100 * time.Millisecond)
  for i:=1;i<5 ;i++{
    fmt.Println(s)
  }
}
