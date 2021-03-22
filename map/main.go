package main

import (
  "fmt"
)

func myFuncMap(k string, v string) (myMap map[string]string){
  myMap = map[string]string{
    k : v,
  }
  return myMap
}

func main(){
  var k string
  var v string
  myMap := make(map[string]string)
  
  fmt.Println("hello! enter with key value: ")
  fmt.Scan(&k)
  fmt.Println("enter with value for key: ")
  fmt.Scan(&v)
  
  myMap = myFuncMap(k,v)
  
  fmt.Println("key: ", k)
  fmt.Println("value: ", v)
  fmt.Println("my map is: ", myMap)
}
  
