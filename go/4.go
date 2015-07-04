package main

import (
    "fmt"
    "strconv"
    "unicode/utf8"
)

func isPalindrome(drome string) bool {

    dromelen := utf8.RuneCountInString(drome)

    // base case
    if dromelen == 1 || dromelen == 0 {
        return true
    }

    // induction
    if drome[0] == drome[dromelen - 1] {
        substring := drome[1:dromelen - 1]
        return isPalindrome(substring)
    }

    return false
}

func main() {

    largest := 0

    for x := 100; x < 1000; x++ {
        for y := 100; y < 1000; y++ {
            z := x * y
            zstr := strconv.Itoa(z)
            if (isPalindrome(zstr) && largest <= z) {
                largest = z
            }
        }
    }

    fmt.Printf("%d\n", largest)
}
