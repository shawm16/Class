package main

import "fmt"

func main(){
  var (
    limit = 4000000
    sum = 0
  )
  /* This is a example of defining and using a array instead
   of 3 different variables you have x[0-2] 0, 1, 2, = 3 variables*/
  var x [3]int
  x[0] = 1
  x[1] = 1
  x[2] = x[0] + x[1]


  /*in Go the for statement is the while statement */
  for x[2] < limit {
    sum += x[2]
    x[0] = x[1] + x[2]
    x[1] = x[2] + x[0]
    x[2] = x[0] + x[1]
  }
  fmt.Println(sum)

}
