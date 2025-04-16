package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
	
}
func main(){
   num1 := 10
   num2 := 20 
   total:= sum(num1,num2)
   fmt.Println(num1,"+",num2,"=",total)	
    
}