package main

import (
  "os"
  "fmt"
  "math"
  "strconv"
)

const Delta = 0.00000001

func Sqrt(x float64) (float64, int) {
  currentGuess := 1.0
  for i := 1; ;i++ {
    nextGuess := guess(x, currentGuess)
    if goodEnough(currentGuess, nextGuess) {
      return currentGuess, i
    } else {
      currentGuess = nextGuess 
    }
  }
}

func guess(arg float64, current float64) float64 {
    return current - ((math.Pow(current,2) - arg) / (2 * current)) 
}

func goodEnough(last float64, current float64) bool {
  return math.Abs(last - current) < Delta
}

func main() {
  arg, err := strconv.ParseFloat(os.Args[1],64)
  if (err == nil) {
    result, iterations := Sqrt(arg)
    fmt.Printf("Iterations: %v\n", iterations)
    fmt.Printf("My Sqrt   = %v\n", result)
    fmt.Printf("math.Sqrt = %v\n", math.Sqrt(arg))
  } else {
    fmt.Printf("Error calculating square root: %v\n",err)
  }
}
