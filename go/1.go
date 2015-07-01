package main

import "fmt"
import "math"

func main() {
    
    acc := 0
    for i := 0; i < 1000; i++ {
        f := float64(i)
        if (math.Mod(f, 3) == 0 || math.Mod(f, 5) == 0) {
            acc += i
        }
    }

    fmt.Printf("%d\n", acc)
}
