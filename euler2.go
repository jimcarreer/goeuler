package main

import "fmt"

func main() {
    var swap = 0
    var num1 = 0
    var num2 = 1
    var sum  = 0
    for num2 < 4000000 {
        if num2 % 2 == 0 {
            sum += num2
        }
        swap = num2
        num2 = num2 + num1
        num1 = swap
        fmt.Printf("Fib num : %d\n",num1)
    }
    fmt.Printf("Sum : %d\n",sum)
}