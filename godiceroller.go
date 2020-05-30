package main

import (
  "fmt"
  "math/rand"
  "flag"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  sides := flag.Int("sides", 6, "number of sides")
  dice := flag.Int("dice", 1, "number of dice")

  flag.Parse()
  if *sides <= 0 || *dice <= 0 {
    fmt.Println()
  }

  fmt.Println("sides:", *sides)
  fmt.Println("dice:", *dice)

  rolls := GenerateRolls(*sides, *dice)
  fmt.Println(rolls)
  fmt.Println(len(rolls))
}

func GenerateRolls( s, d int ) []int {
  rolls := []int{}
  for i := 1; i <= d; i++ {
    roll := generateRandom(s)
    rolls = append(rolls, roll)
  }
  return rolls
}

func generateRandom( n int ) int {
  return rand.Intn(n)+1
}
