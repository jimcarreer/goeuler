package main

import "fmt"
import "strconv"

// 999*999 = 998001 (largest product of two 3 digit #)
// 100*100 = 10000 (smallest product of two 3 digit #)
// smallest palindromatic 10001  over  10000
// largest  palindormatic 997799 under 998001

func main() {
    fmt.Printf("Answer : %d\n", FindPal())
}

func FindPal() int {
    var itop = 997
    var fac1 = 100
    var fac2 = 100
    var ipal = 997799
    for itop > 99 {
        ipal = MakePal(itop)
        fmt.Printf("Pal : %d\n", ipal)
        itop -= 1
        fac1 = 100
        for fac1 < 999 {
            fac2 = ipal / fac1
            if ipal % fac1 == 0 {
                if fac2 >= 100 && fac2 <= 999 {
                    return ipal
                }
            }
            fac1 += 1
        }
    }
    return 0
}

func MakePal(tophalf int) int {
    var s = strconv.Itoa(tophalf)
    var runes = []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    var i , _ = strconv.Atoi(s+string(runes))
    return i
}