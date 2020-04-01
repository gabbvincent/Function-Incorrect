package main

import (
  "fmt"
  "math/rand"
  "time"
)

func incorrect(){
  var b [5]string
  b[0] = "Incorrect!"
  b[1] = "Wrong!"
  b[2] = "Nope!"
  b[3] = "That's not it!"
  b[4] = "Try again!"

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  i := r1.Intn(5)

  fmt.Println(b[i])
}

func main() {
  
  for i := 0; i < 10; i++{
    incorrect()
  }
}