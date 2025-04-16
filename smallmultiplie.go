package main 
import "fmt"


func smallestEvenMultiple(n int) int {
    if n%2==0{
        return n 
    }else{
        return n*2
    }
}
func main(){
	num:= 6
	fmt.Println(smallestEvenMultiple(num))
}