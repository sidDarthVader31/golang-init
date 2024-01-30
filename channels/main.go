package main

import "fmt"

func main(){

  //declare a channel 
  var ch chan int 
  fmt.Println(ch)

  ch = make(chan int) //initialize a channel 
  fmt.Println(ch) //it holds an address 

  //create and initialize a channel
  ch1 := make(chan int)



  //close a channel 
   close(ch)

  //create a bidirectional channel 
  chBi := make(chan int)

  //create a unidirectional channel that only receives data 
  chReceive := make(<- chan string)

  //create a unidirectional channel that only sends data 
  chSend := make(chan<- string )

  fmt.Printf("%T, %T, %T\n", chBi, chReceive, chSend)

  go send(10, ch1)
  nums := <-ch1
  fmt.Println("value received is :", nums)
  fmt.Println("exiting main function")
}

func send(n int, ch chan int){
  ch <-n  //send data in channel 
}
