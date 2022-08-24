package main

import "fmt"


func main() {
// printing number from 1 - 10

for i := 1; i <= 10; i++ {

  fmt.Println(i)

}
  fmt.Println("End")

// using range to print array

var array = []int { 1,3,4,5,6,7,8,9,0}

for _, num := range array {
  fmt.Println(num)
}

//Using if statement
var num = 2
if num == 2 {
  fmt.Printf("We got it")
} else{
  fmt.Printf("Got it wrong")
}
// while loop part
 whilenum := 0
for whilenum < 10 {

  fmt.Println(whilenum)
  whilenum ++
}

}
