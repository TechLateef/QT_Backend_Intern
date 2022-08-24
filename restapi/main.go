
package main

import(
  "fmt"
  "log"
  "net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "Welcome to QTSolution intern! ")
  fmt.Println("Endpoint hit: homepage")

}

func Name(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "My names are: Mubarak Abdullateef ")
  fmt.Println("Endpoint hit: homepage")

}

func Stack(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "My stack: Backend Dev ")
  fmt.Println("Endpoint hit: homepage")

}

func handleRequest(){
http.HandleFunc("/", Welcome)
http.HandleFunc("/Name", Name)
http.HandleFunc("/Stack", Stack)
log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
handleRequest()

}
