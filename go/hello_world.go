package main
import (
    "fmt"
    "strings"
)

func main() {
    var inputString [6]string
    fmt.Scan(&inputString[0])
    fmt.Scan(&inputString[1])
    fmt.Scan(&inputString[2])
    fmt.Scan(&inputString[3])
    fmt.Scan(&inputString[4])
    fmt.Scan(&inputString[5])
    fmt.Println("Hello, World.")
    fmt.Println(strings.Join(inputString[:], " "))
}
