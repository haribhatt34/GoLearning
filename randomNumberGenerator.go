package main

// With go we have to type fully qualified path for the particular sub package,
// like Intn is in rand subpackage, so we have to type path till there in import.
import ("fmt"
	"math/rand")

func calc() {
    fmt.Println("A random generator number:  ",rand.Intn(100))
}

func main() {
    calc()
}
