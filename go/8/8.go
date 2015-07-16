package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

// Reads each of the lines in a file and then
// outputs all of the lines as a slice of strings. 
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines, scanner.Err()
}

// Builds an array of integers from an
// array of strings of integers.
func buildValueArray(scannerInput []string) []int {
    valueString := strings.Join(scannerInput, "")

    var valueArray []int

    for _,character := range valueString {
        // Note here that the (- '0') trick only works
        // for the integers 0 to 9 for conversion
        // to integers in the type cast.
        value := int(character - '0')
        valueArray = append(valueArray, value)
    }
    return valueArray

}

func main() {
    // Read the lines from the text file with the values.
    lines, err := readLines("number.txt")
    // Error handling for the readline function.
    if err != nil {
        fmt.Println(err)
    }

    // From the lines in the previous data,
    // build a slice of integers.
    valueArray := buildValueArray(lines)

    // Finally, determine the largest product.
    largestProduct := 0

    for i := 0; i < len(valueArray) - 12; i++ {
        currentProduct := 1

        for j := 0; j < 13; j++ {
            currentProduct *= valueArray[i + j]
        }

        if currentProduct > largestProduct {
            largestProduct = currentProduct
        }
    }

    fmt.Printf("%d\n", largestProduct)
}
