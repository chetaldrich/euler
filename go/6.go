package main

import (
    "fmt"
)

func sumSquares(value int) int {
    acc := 0
    for i := 1; i <= value; i++ {
        acc += i * i
    }
    return acc
}

func squareSum(value int) int {
    acc := 0
    for i := 1; i <= value; i++ {
        acc += i
    }
    acc = acc * acc
    return acc
}

func main() {
    result := squareSum(100) - sumSquares(100)
    fmt.Printf("%d\n", result)
}
