//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
package main
import "fmt"

// I didn't write a program for this, the answer is pretty obvious
// Consider :
// 1 2 3 (2 2) 5 (2 3) 7 (2 2 2) (3 3) (2 5)
// 2520 factored:
// 2 2 2 3 3 5 7
// We merely need to find the lowest power of each prime <= 20
// that is highest power of a prime for every composite <= 20:
// 1 2 3 (2 2) 5 (2 3) 7 (2 2 2) (3 3) (2 5) 11 (2 2 3) 13 (2 7) (3 5) (2 2 2 2) 17 (2 9) (2 2 5)
// For 2 = 2 2 2 2 (16)
// For 3 = 3 3 (9)
// All the rest either contain different primes of power 1
// A general solution to the problem would require an input X being the maximum factor, factoring each number 1 ... X to prime lists
// and determining the maximum power of each prime under X iteraviely by counting their appearace in the list


func main() {
    var primes = []int{  2,  3,  5,  7, 11, 13, 17, 19}
    var powers = []int{  4,  2,  1,  1,  1,  1,  1,  1}
    fmt.Printf("Answer : %d\n", Product(primes, powers))
}

func Product(primes []int, powers []int) int {
    var product = 1
    for i := 0; i < len(primes); i++ {
        for j := 0; j < powers[i]; j++ {
            product *= primes[i]
        }
    }
    return product
}
