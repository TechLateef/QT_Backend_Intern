

package main
import (
  "fmt"
  "os"
  "bufio"
)

func main() {
// Collecting input from user
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter your name \n")
name, _ := reader.ReadString('\n')

fmt.Println("Your name is: ", name)


}
