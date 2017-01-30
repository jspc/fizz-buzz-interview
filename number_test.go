package main

import (
    "testing"
)

func TestNumberToString(t *testing.T) {
    for _,test := range []struct{
        Name string
        Value int
        Result string
    }{
        {"Modulo 3", 3, "fizz"},
        {"Modulo 5", 5, "buzz"},
        {"Both Modulo 3 and Modulo 5", 15, "fizzbuzz"},
        {"Neither Modulo 3 nor Modulo 5", 7, "7"},
    }{
        t.Run(test.Name, func (t *testing.T) {
            n := Number{test.Value}
            v := n.ToString()
            if v != test.Result {
                t.Errorf("%s error: expected %q, received %q", test.Name, test.Result, v)
            }
        })
    }

}
