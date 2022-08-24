package main

import (
  "fmt"
)

func main() {
  // array of type int
var array = []int {1,2,3,4,5}
// array of type  string
var arraystring = []string {"Mubarak", "Backend dev" , "Qt-intern-program"}
// array slice
arrayslice := arraystring[0:2]

fmt.Println("your numbers are: \n", array)
//printing array with the other print Printf()
fmt.Printf("YOur informations are as follow: %s \n ", arraystring)

fmt.Print("Using array slice to print Specific items \n ", arrayslice)

}
