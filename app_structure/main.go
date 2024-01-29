// a package clause starts every source file main is a special name declaring an executable rather than a library package package main

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode/utf8"
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
  fmt.Println("hey")
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

  //slices with arrays 
  arr := [4] int {10,20,30,40}
  sarr1, sarr2 := arr[0:2], arr[1:3]
  fmt.Println(arr, sarr1, sarr2)
  //modifying arr 
  arr[1] = 3 // this will modify the slices as well since arr is their backing array 
  fmt.Println(arr, sarr1, sarr2);
  sarr1[1] = 1000 // this wil also modify the slices and array arr 
  fmt.Println("modifying sarr1")
  fmt.Println(arr, sarr1, sarr2)

  // if we want to copy a slice from another slice we can use append if we do not want slices to share 
  // the same backing array 

  cars := []string {"Tata","Mahindra", "Kia", "Hyundai"}
  newCars := []string {}
  newCars= append(newCars, cars...)
  fmt.Println("cars:", cars, newCars)

  cars [0] = "Honda" // this will not modify newCars 
  fmt.Println(cars, newCars)

  //stringsk bytes and runes

    // characters or rune literals are expressed in Go by enclosing them in single quotes
    // declaring a variable of type rune (alias to int32)
    var1 := 'a'
    fmt.Printf("Type: %T, Value:%d \n", var1, var1) // => Type: int32, Value:97
    // value is 97 (the code point for a)
 
    // declaring a string value that contains non-ascii characters
    str := "ţară" // ţară means country in Romanian
    // 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.
 
    //The len() built-in function returns the no. of bytes not runes or chars.
    fmt.Println(len(str)) // -> 6,  4 runes in the string but the length is 6
 
    // returning the number of runes in the string
    m := utf8.RuneCountInString(str)
    fmt.Println(m) // => 4
 
    // by using indexes we get the byte at that positioin, not rune.
    fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163
 
    // decoding a string byte by byte
    for i := 0; i < len(str); i++ {
        fmt.Printf("%c", str[i]) // -> Å£arÄ
    }
 
    fmt.Println("\n ####")
 
    // decoding a string rune by rune manually:
    for i := 0; i < len(str); {
        r, size := utf8.DecodeRuneInString(str[i:]) // it returns the rune in string in variable r
        //and its size in bytes in variable size
 
        // printing out each rune
        fmt.Printf("%c", r) // -> ţară
        i += size           // incrementing i by the size of the rune in bytes
    }
 
    fmt.Println("\n#####")
 
    // decoding a string rune by rune automatically:
    for i, r := range str { //the first value returned by range is the index of the byte in string where rune starts
        fmt.Printf("%d -> %c\n", i, r) // => ţară
    }

  str1 := "I !love golang !"
  str2 := "ðß©˳þ"
  fmt.Println(str1[0:5]) // slicing string with ascii values work normally
  fmt.Println(str2[0:3]) // slicing strings with non ascii values will return abrupt strings
  //because slicing happens for bytes not runes 

  //to properly slice the string (contatining non ascii values) we will have to convert it to rune

  rune2 := []rune(str2)
  fmt.Printf("%T\n", rune2)

  fmt.Println(string(rune2[0:3]))

  //maps

  var checkMap map[string] int //declare an non initialized map 
  fmt.Println("check map:", checkMap) 

  //checkMap["sid"] = 1; // this is an invalid line, we cannot assign a key value pair to a map if its not initialized 

  peopleMap := map[string] int{} //declaring an initialized map 
  peopleMap["sid"]=1
  fmt.Println("check map :", peopleMap)

  map1 := make(map[string]int)
  map1["sid"] = 345
  fmt.Println("map1:", map1)

  //fetching a non existing key returns the zero value of the type of value for the map 
  val1 := peopleMap["sid"]
  fmt.Println(val1) //returns 1
  val2 := peopleMap["siddharth"] //val2 is zero 
  fmt.Println(val2)

  val3, ok := peopleMap["sid"] // this returns two values, val3 = value at key and ok is boolean 
  // which is true if key exists and false if key does not exist

  fmt.Printf("%d : %v\n", val3, ok)

  //delete a key from map 

  delete(peopleMap, "sid")
  fmt.Println(peopleMap)

  nameMap :=  map[string] string {
    "sid":"sid",
    "siddharth":"siddharth",
  }
  nameMap1 := nameMap //this creates a copy with same map header, any change to either one of the map
  // will be reflected back in the other as well 
  nameMap1["sid"] ="sid1"
  fmt.Println(nameMap)
  
  //structs

  type casting struct {
    name string 
    gender string 
  }
  type movie struct{
    title string
    director string
    year int
    cast casting
  }
  movie1 := movie{"Interstellar", "Christopher Nolan", 2014, casting{"Anne hatheway", "Female"}}
  fmt.Println(movie1)

  // struct fields can be accessed by . operator 

  fmt.Printf("director of movie %v is %v\n", movie1.title, movie1.director)
  fmt.Printf("cast of movie %v is : %v\n",movie1.title, movie1.cast.name)

  //functions
  f11()
  f2(10,20)
  f3(4,5,6,5.5,6.5,"golang")
  sum := f4(5, 5)
  fmt.Println(sum)
  sum1, multiple1 := fMultiple(10, 20)
  fmt.Println(sum1, multiple1)

  //variadic functions
  fVariadic(1,2,3,4,5,6)

  numSlice := []int{1,2,3,4}
  numSlice = append(numSlice, 5,6)
  fmt.Println("nusslice:", numSlice)
  fVariadic(numSlice...)

  //pointers details 

  var x3 int = 5 
  pointerA := &x3
  fmt.Printf("type of pointerA : %T, and value is : %v\n", pointerA, pointerA)

  var x4 int =50
  fPointers(&x4)
  fmt.Println("x4 value is:", x4)

  //OOPS

  //receiver functions 

  // we can associate a function to a type with the help of receiver functions
  //function declaration for this type of name is given in printNames function

  peopleNames := names{"Luka", "Vini", "toni"}
  peopleNames.printNames()

  //we can also invoke printNames on the type 
  names.printNames(peopleNames)

  myCar := car{brand: "Tata", price: 2000}
  fmt.Println(myCar)

  //changing mycar properties
  myCar.changeCar("Audi", 40000)
  fmt.Println("changed my car config:", myCar)

  //interfaces 

  rect := rectangle{width: 10., height: 20.}
  cir := circle{radius: 20.}
  printShape(rect)
  printShape(cir)
  circleArea := cir.area()
  fmt.Println("circle area:", circleArea)

  var sShape shape
  fmt.Printf("type of s shape is : %T\n", sShape)

  sShape = cir
  fmt.Printf("type of s shape now :%T\n", sShape)

  //empty interface
  var empty emptyInterface
  empty = 5
  fmt.Println(empty)
  empty = "name"
  fmt.Println(empty)

  //embedded interfaces
  c := cube{edge: 5., color: "blue"}
  display(c)
}


type emptyInterface interface {}

type cube struct {
  edge float64
  color string
}
func (c cube) getColor() string {
  return c.color
}
func (c cube) area() float64{
  return 6 * math.Pow(c.edge, 2)
}
func (c cube) perimeter() float64{
  return 12 * c.edge
}
func (c cube) volume() float64{
  return math.Pow(c.edge,3)
}
func display(g geometry){
  fmt.Println("area :", g.area())
  fmt.Println("perimeter:", g.perimeter())
  fmt.Println("volume:", g.volume())
  fmt.Println("color:", g.getColor())
} 
type shape interface{
  area() float64
  perimeter() float64
}

type rectangle struct{
  height float64
  width float64
}


type object interface {
  volume() float64
}

type geometry interface{
  shape
  object
  getColor() string
}
func (r rectangle) area() float64{
  return r.width * r.height
}
func (c circle) area() float64{
  return math.Pi * math.Pow(c.radius,2 )
}
type circle struct {
  radius float64 
}

func (r rectangle) perimeter() float64{
  return 2 * r.height * r.width
}
func (c circle) perimeter() float64{
  return 2 * math.Pi * c.radius
}

func printShape(s shape){
  fmt.Printf("shape: %#v\n", s)
  fmt.Println("area:", s.area())
  fmt.Println("perimeter:", s.perimeter())
}

type names [] string
type car struct{
  brand string
  price int
}
//receiver functions
func (n names)printNames(){
  for i,name := range n {
    fmt.Println(i, name);
  }
}

//receiver function with pointer 

func (car *car) changeCar(brand string, price int){
  (*car).price = price
  (*car).brand = brand
}


func fPointers(a *int){
  *a +=10
}
func fVariadic(a... int){
  fmt.Printf("%T\n", a )
  fmt.Printf("%#v\n", a )
}
func f11(){
  fmt.Println("called form f1")
}

func f2(a int, b float64){
  fmt.Println(a, b)
}

//shorthand declaration

func f3(a,b,c int, d,e float64, s string){
  fmt.Println(a,b,c,d,e,s);
}

//return types 

func f4(a, b int) (s int) {
  s = a+b
  return 
}

//multiple return types 


func fMultiple(a, b int) (int, int){
  return a+b, a*b;
}

