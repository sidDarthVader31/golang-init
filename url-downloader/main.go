package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main(){
  url := [] string{"https://www.google.com", "https://www.youtube.com", "https://www.gmail.com"}
  var wg sync.WaitGroup
  wg.Add(len(url))
  for _,v := range url{
    go checkUrl(v, &wg)
    fmt.Println(strings.Repeat("#", 10))
  }
  wg.Wait()
}

func checkUrl(url string, wg *sync.WaitGroup){
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

      fmt.Println("writing the html to a file for url:", url)
      err =  os.WriteFile(fileName, bodybytes, 0664)
      if err != nil{
        fmt.Println("error while writing file for ", err)
      }else{
        fmt.Println("file written successfully for url", url)
      }
    }else{
      fmt.Println("status code is not 200", resp.StatusCode)
    }
  }
  wg.Done()
}
