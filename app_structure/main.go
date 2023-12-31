// a package clause starts every source file main is a special name declaring an executable rather than a library package package main
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const secondsInhour = 3600
type km float64
type miles float64
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

  //arrays
  var numbers = [4] int {1,2,3,4}
  fmt.Printf("numbers array : %T \n",numbers)

  //slices
  var slice =[]string{"london","paris","milan","madrid"}
  fmt.Printf("slice:%T \n", slice)


  //maps
  data := map[string]float64{
    "key1": 1.4,
    "key2": 2.4,
  }
  fmt.Printf("maptype: %T\n", data);

  //struct

  type person struct{
    name string
    age int
    designation string
  }
  //var newUser Person;
  //fmt.Printf("%T", newUser);
  
  //pointers
  var x int = 2
  var xAddress *int = &x
  fmt.Printf("type of xAddress: %T , and value is %v \n", xAddress, xAddress);

  var f float32 = math.MaxFloat32
    fmt.Println(f)
  f = f * 1.5
  fmt.Println("new f:", f)

  //type conversions

  var x1 = 3
  var x2 = 5.2
  fmt.Println(float64(x1) * x2)
  fmt.Println(int(x2) * x1)


  //CONVERt string to int 

  iString, err := strconv.Atoi("45");
  s := strconv.Itoa(-42)
  b, errb := strconv.ParseBool("true")
  f1, err := strconv.ParseFloat("3.1415", 64)
  i, err := strconv.ParseInt("-42", 10, 64)
  u, err := strconv.ParseUint("42", 10, 64)
  fmt.Printf("%T %T %T %T %T %T \n", iString, b, s, i, u, f1)
  _ = err
  _ = errb

  //defined types 

  var distanceKm km = 475
  var distanceInMiles miles 
  distanceInMiles = miles(distanceKm)/ 1.6 // type conversion of defined types
  fmt.Println("distance in miles:", distanceInMiles)


  fmt.Println("command line args:", os.Args)

}
