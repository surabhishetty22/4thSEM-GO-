package main
import (
    "fmt"
    "math"
)
func largestprimefactor(n int) int {
    //initialize ther largest prime factor as 1
    largest := 1
    //divide the number  by 2 repeatedly if it is divisible by 2
    for n%2 ==0 {
        largest = 2
        n /= 2
    }
    //check for divisiblityby odd numbers starting from 3 upto square root of n
    for i := 3; i<= int(math.Sqrt(float64(n))); i+=2 {
        for n%i == 0 {
            largest = i
            n /= i
        }
    }
    //if n is a prime no greater than 2 ,it is largest prime factor
    if n > 2 {
        largest = n
    }
    return largest
}
func main() {
    number := 14
    largest := largestprimefactor(number)
 
  fmt.Println("the largest prime factor of %d is %d \n", number, largest)
}