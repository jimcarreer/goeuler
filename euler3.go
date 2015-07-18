package main

import "fmt"

func main() {
    var what = 600851475143
    var divisor = 2
    var largest = 0
    for what > 1 {
        if what % divisor == 0 {
            what = what / divisor
            divisor = 2
        } else {
            divisor = divisor + 1
            if divisor > largest {
                largest = divisor
            }
        }
    }
    fmt.Printf("Answer : %d\n", largest)
}

