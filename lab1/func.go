package main
import "fmt"
//task1
func add(a int, b int) int {
    return a + b
}
//task2
func swap(s1, s2 string) (string, string) {
    return s2, s1
}
//task3
func quotientAndRemainder(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    sum := add(10, 20)
    fmt.Println("Sum:", sum)  

    str1, str2 := swap("first", "second")
    fmt.Println("Swapped:", str1, str2)  

    quotient, remainder := quotientAndRemainder(20, 3)
    fmt.Println("Quotient:", quotient, "Remainder:", remainder) 
}
