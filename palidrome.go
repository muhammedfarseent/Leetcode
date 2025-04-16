package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
    return false
    }
	front := x
	back := 0
	for x > 0 {
		lastDigit := x % 10
		back = back*10 + lastDigit
		x /= 10
	}
	return front == back
}
func main(){
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1256))
	fmt.Println(isPalindrome(-121))
}