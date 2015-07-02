package main

import "fmt"
import "math"

func main() {
    acc := 0
    fib1 := 1
    fib2 := 2
    tempfib := 1

    for fib2 < 4000000 {
        f := float64(fib2)
        if (math.Mod(f, 2) == 0) {
            acc += fib2
        }

        tempfib = fib2 + fib1
        fib1 = fib2
        fib2 = tempfib
    }

    fmt.Printf("%d\n", acc)
}
