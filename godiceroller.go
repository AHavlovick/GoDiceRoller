package main

import (
  "fmt"
  "math/rand"
  "os"
  "strconv"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  n := os.Args[1]
  num, err := strconv.Atoi(n)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(generateRandom(num))
}

func generateRandom( n int) int {
  return rand.Intn(n)+1
}
