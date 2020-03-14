package main

import "fmt"


//func add (a, b float64) float64 {

func add (a, b float64) float64 {
    return a + b
}

func printMe (a, b string) (string, string) {
    return a, b
}


func main() {
//  var a float64 = 5.6
//  var b float64 = 4.4
//  var a, b float64 = 5.6, 4.4
//  type for below deducted during run time.
    a, b := 5.6, 4.4
    fmt.Println(add(a, b))

    str1, str2 := "hey", "hari"
    fmt.Println(printMe (str1, str2))

    var d float64 = 45.5
    var c int = int (d)
    m := c // m will be of type int

    fmt.Println(d, c, m)
}
