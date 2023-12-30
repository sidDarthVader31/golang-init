// a package clause starts every source file
// main is a special name declaring an executable rather than a library package
package main
import "fmt"
const secondsInhour = 3600

//main is the special function name ans is used as the main entrypoint of the executable file.
// main function is mandatory in an executable file
func main(){
  var a int = 10
  fmt.Println("hello world", a)
  distance := 60.8
  fmt.Println("the distance in miles is :", distance * 0.6)
}
