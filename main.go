package main

import (
    "flag"
    "fmt"
)

var (
    rangeVar int
)

func init() {
    flag.IntVar(&rangeVar, "range", 300, "Range of values to run fizz buzz against")
    flag.IntVar(&fizz, "fizz", 3, "Value of fizz; multiples of which will be replaced with 'fizz'")
    flag.IntVar(&buzz, "buzz", 5, "Value of buzz; multiples of which will be replaced with 'buzz'")
}

func main() {
    flag.Parse()

    for v := range make([]int, rangeVar) {
        n := Number{v}
        fmt.Printf("%q ", n.ToString())
    }
}
