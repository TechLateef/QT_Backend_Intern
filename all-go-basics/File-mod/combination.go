
// this file compose of everything learn from file module
package main
import (
  "fmt"
  "os"
    // "bufio"
    // "io /ioutil"
)


func reader() {

  // the part read the file

  b, err := os.ReadFile("content.txt")

  if err != nil {

    fmt.Print(err)

  }
  // this convert the file read into a string
  var result = string(b)

  fmt.Println("File found and this what's inside: ", result)
}


func main(){
  // this part create a file
file, err := os.Create("content.txt")
if err != nil {
  fmt.Print(err)
}
defer file.Close()

file.WriteString("Welcome to go")

// this will check if it exist if yes it will read

if _, err := os.Stat("content.txt");
err == nil {

reader()

} else{

  fmt.Print("File cant be found")

}




}
