package main 

import (
	"fmt"
	"strings"
)
// function for to converting 
func ToLower(s string)string{
	return strings.ToLower(s)
}

func main(){
	// declaring and changing of this 
	orginal:="HELLO"
	changed:=ToLower(orginal)
	// printing the both 
	fmt.Println("this is orginal one ",orginal)
	fmt.Println("this is changed one ",changed)
}
// we can add with our convinence 