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

  //if else 

  i1, err1 := strconv.Atoi("45a");
  if err1 != nil {
    fmt.Println(err1)
  } else {
    fmt.Println(i1)
  }

  //simpleif


  if i5,err5 := strconv.Atoi("45sd"); err5 == nil {
    fmt.Println("no error, value is ", i5)
  } else{
    fmt.Println("error found:", err5)
  }


  if num := 10; num > 5 {
    fmt.Println("number greater than 5")
  }else{
    fmt.Println("number less than 5 ")
  }

//for loop

  for i:=0;i<10;i++{
    fmt.Println(i)
  }

  //for loop over an array 

  people := [5] string {"Jude","Vini","Rodrgyo","Luca","Toni"};

  for index, name := range people{
    fmt.Printf("name at index: %d is %s\n", index, name) 
  }

  //arrays 
  nameArray := [4]string{"sid","sid2","sid3"}
  fmt.Printf("%#v\n", nameArray)


  //epplisis operator
  numArray := [...]int {1,2,3,4,5}
  fmt.Printf("%#v\n", numArray)
  fmt.Printf("length of num array is : %d\n", len(numArray))

  //iterate over an array 

  for i,v := range(numArray){
    fmt.Printf("index: %d, value: %d\n", i,v)
  }

  //2d arrays 

  twoDArray := [2][3]int{
    {5,6,7},
    {8,9,10},
  }
  fmt.Println(twoDArray)

  mArray := [3] int {1,2,3}
  nArray :=mArray;

  //above line will create a new copy of mArray , any modification to mArray will not reflect back on nArray
  ///however in slices the two slices will be connected and same copy will be used 

  fmt.Printf("address of mArray: %p\n", &mArray)
  fmt.Printf("address of n Array:%p \n", &nArray);
  //the above two addresses will be different

  //keyed arrays 

  keyedArray := [3]int {
    1:10,
    2: 20,
    0:5, // WHILE declaring multi line array , it is mandatory to add a , at the end of last line
  }
fmt.Println(keyedArray)
  keyedNames := [...]string{
    4:"sid4",
    3: "sid3",
  }
  fmt.Println(keyedNames)
  fmt.Println("length of keyed names:", len(keyedNames))


  //not specifiying the index of an element 

  keyedCities := [...] string {
    4: "London",
    "New Delhi",
    2: "Dehradun",
  }
  fmt.Printf("%#v\n", keyedCities)
  //unindexed value will take the index after the index of the value above it, in this case london 
  // if no value is specified before unindexed value, meaning the first value we add is of unindexed
  //type then it will take the index 0

  //slices
  var cities[] string
  fmt.Println("slice:", cities== nil) // nil means the value exist but hasnt been initialized yet

  fmt.Println("length of cities:", len(cities))

  numbersSlice := []int {1,2,3,4}
  fmt.Println("numbers::slice:", numbersSlice)

  //creating a slice using make 

  numsSliceMake := make([]int, 2)
  
  fmt.Printf("using make :%#v", numsSliceMake)

  var slice1 []int; //equal to nil
  slice2 := []int{};//not equal to nil

  fmt.Println(slice1 == nil) //returns true
  fmt.Println(slice2 ==  nil) //returns false 

  //note : in go slices can only be compared with nil
  //to compare two slices we will have to iterate over them and compare value at every index

  //append to a slice 
  fmt.Println(numbersSlice) //prints - 1,2,3 ,4 
  numbersSlice = append(numbersSlice, 5)

  fmt.Println("numbers slice after append:", numbersSlice) // prints 1,2,3,4,5
  
  //append can also be used to add a slice to another slice 
  nSlice :=[]int {40,50,60}
  numbersSlice = append(numbersSlice, nSlice...)
  fmt.Println("numbers slice after adding another aslice:", numbersSlice)

  s1 := [] int {10,20,30,40,50,60}
  s3,s4 := s1[0:2], s1[1:3] // s1, s3 and s4 all contains same backing array 
  fmt.Println(s3)
  fmt.Println(s4)

  s3[1]= 1000 // when we modify s3 the backing array is also modified which changes the value of s1 and s4 too 
  fmt.Println(s4) // prints [1000,30]
  fmt.Println(s1) // prints [10,1000,30,40,50,60]







}
