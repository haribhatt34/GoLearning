package main

import "fmt"

func main() {

    a := 5
    b := &a
    c := &b
    fmt.Println("a: ", a, ", b: ", b, ", *b: ", *b, ", **c: ", **c)
}
