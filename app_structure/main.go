// a package clause starts every source file
// main is a special name declaring an executable rather than a library package
package main

import "fmt"

const secondsInhour = 3600

// main is the special function name ans is used as the main entrypoint of the executable file.
// main function is mandatory in an executable file
func main() {
	var a int = 10
	fmt.Println("hello world", a)
	distance := 60.8
	fmt.Println("the distance in miles is :", distance*0.6)
  
  figure := "Circle"
  radius := 10
  pi := 3.14

  fmt.Printf("figure is %s, radius is %d, area is %f \n", figure, radius, pi*float64(radius)*float64(radius))
  const (min1 = 100
    min2 =200
    min3 ="sid" )
  fmt.Println(min1, min2, min3)
}
