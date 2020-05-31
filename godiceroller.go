package main

import (
  "fmt"
  "math/rand"
  "flag"
  "time"
  "os"
)

type CustomError struct {
  Arg string
  Val int
  Msg string
}

type UserInput struct {
  Sides int
  Dice int
  Errors []error
  Rolls []int
}

// Custom error handling
func (e *CustomError) Error() string {
  return fmt.Sprintf("Error: %s-->%s:<%d>", e.Msg, e.Arg, e.Val)
}

// Set user Input to Struct
func InitializeUserInput(s, d int) *UserInput {
    return &UserInput{
      Sides: s,
      Dice: d,
      Errors: ValidateInputs(s, d),
    }
}

func ValidateInputs(s, d int) []error {
  errors := []error{}
  intErrMsg := "Values must be positive integers"

  if s <= 0 {
    sideerr := &CustomError{"sides", s, intErrMsg}
    errors = append(errors, sideerr)
  }

  if d <= 0 {
    diceerr := &CustomError{"dice", d, intErrMsg}
    errors = append(errors, diceerr)
  }
  if errors != nil {
    return errors
  }
  return nil
}

func GenerateRolls( s, d int ) []int {
  rolls := []int{}
  for i := 1; i <= d; i++ {
    roll := rand.Intn(s)+1
    rolls = append(rolls, roll)
  }
  return rolls
}

func rollSums( rol []int ) int {
  sum := 0
  for _, val := range rol {
    sum += val
  }
  return sum
}

func main() {
  rand.Seed(time.Now().UnixNano())
  sides := flag.Int("sides", 6, "number of sides")
  dice := flag.Int("dice", 1, "number of dice")
  flag.Parse()

  userInputs := InitializeUserInput(*sides, *dice)
  if len(userInputs.Errors) > 0 {
    for _, errval := range userInputs.Errors {
      fmt.Println(errval)
    }
    os.Exit(0)
  }
  rolls := GenerateRolls(userInputs.Sides, userInputs.Dice)
  rollsum := fmt.Sprintf("Sums: %d", rollSums(rolls))
  fmt.Println("Rolls:", rolls)
  fmt.Println(rollsum)
}
