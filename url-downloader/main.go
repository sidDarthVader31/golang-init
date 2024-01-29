package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main(){
  fmt.Println("args length/:", len(os.Args))
  checkUrl(os.Args[1]) 
}

func checkUrl(url string){
  fmt.Println("url is :", url)
  resp, err := http.Get(url)
  if err != nil{
    fmt.Println(err)
    fmt.Printf("%s is down!!\n", url)
  }else{
    defer resp.Body.Close()
    fmt.Printf("%s's status code is : %d \n",url,  resp.StatusCode)
    if resp.StatusCode == 200 {
      bodybytes, err :=  ioutil.ReadAll(resp.Body);
      fileName := strings.Split(url, "//")[1]
      fileName = fileName + ".txt"
      fmt.Println("filename:", fileName)


      //write the contents to a file 

      fmt.Println("writing the html to a file ")
      err =  os.WriteFile(fileName, bodybytes, 0664)
      if err != nil{
        fmt.Println("error while writing file:", err)
      }else{
        fmt.Println("file written successfully")
      }
    }else{
      fmt.Println("status code is not 200", resp.StatusCode)
    }
  }
}
