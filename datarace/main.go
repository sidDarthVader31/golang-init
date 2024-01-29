package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
  var n int = 0
  var wg sync.WaitGroup
  gr := 100;
  wg.Add(gr*2)

  for i:=0;i<100;i++{
    go func() {
      time.Sleep(time.Second /10)
      n++
      wg.Done()
    }()
    go func() {
      time.Sleep(time.Second /10)
      n--;
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println("final value of n is:", n)

}
