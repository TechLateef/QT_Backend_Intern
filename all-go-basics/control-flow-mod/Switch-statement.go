package main

import "fmt"


func main(){
switch x := 2; {
case x == 2:
  fmt.Println("The integer is correct ")
case x == 3:
fmt.Println("You got it wrong")
}

}
