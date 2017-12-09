package main

import "fmt"
import "math"

const name string = "constant"

func main() {
    fmt.Println(name)

    const number = 500000
    fmt.Println(math.Sin(number))

    var a string = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(!d)

    // The := syntax is shorthand for declaring and initializing a variable
    f := "Short"
    fmt.Println(f)    
}