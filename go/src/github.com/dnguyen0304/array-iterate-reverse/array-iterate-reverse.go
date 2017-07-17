package main


// Use the go doc utility to get an object's documentation.
import "fmt"

func createRange(from int, to int) []int {
    length := to - from
    x := make([]int, length)
    // The range keyword in Go is analogous to the enumerate function in
    // Python. Iterating over an empty array or slice (i.e. one of length 0),
    // does not do anything.
    for index, _ := range x {
        x[index] = from + index
    }
    return x
}


func main() {
    // This is a half-open range as with slicing. In other words, it is,
    // inclusive only of the left value.
    x := createRange(0, 10)
    for i := len(x) - 1; i >= 0; i-- {
        fmt.Println(x[i])
    }
}
