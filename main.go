package main

import "fmt"

func main()  {
	var size int
  fmt.Scan(&size)

  var slice []int
  var aux int

  for  i:=0; i<size; i++ {
    fmt.Scan(&aux)
    slice = append(slice, aux)
    //fmt.Println(slice)
  }
  suma := 0
  for _, v:= range slice {
    suma += v
  }
  fmt.Println(suma)
}