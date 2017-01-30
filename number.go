package main

import (
    "fmt"
)

var (
    fizz = 3
    buzz = 5
)

type Number struct {
    Value int
}

func(n Number)ToString() (s string) {
    switch {
    case n.Value % fizz ==0  && n.Value % buzz == 0:
        s = "fizzbuzz"
    case n.Value % fizz == 0:
        s = "fizz"
    case n.Value % buzz == 0:
        s = "buzz"
    default:
        s = fmt.Sprintf("%d", n.Value)
    }

    return
}
