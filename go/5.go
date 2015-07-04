package main

import (
    "fmt"
    "math"
)

func gcd(x, y int) int {
    // courtesy of Euclid
    a := float64(x)
    b := float64(y)

    for b != 0 {
        a, b = b, math.Mod(a, b)
    }

    return int(a)
}

func lcm(x, y int) int {
    return int(math.Floor(float64((x * y) / gcd(x , y))))
}

func main() {

    init := lcm(1, 2)

    for i := 3; i <= 20; i++ {
        init = lcm(i, init)
    }

    fmt.Printf("%d\n", init)
}
