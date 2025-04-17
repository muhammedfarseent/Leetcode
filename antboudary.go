package main 

import "fmt"

func returnToBoundaryCount(nums []int) int {
    position:=0
    count:=0
    for i:=0; i<len(nums);i++{
        position =+nums[i]
        if position==0{
            count++
        }
    }
    return count
}

func main(){
	arr:= [] int{2, -2, 3, -3, 1}
	result:=returnToBoundaryCount(arr)
	fmt.Println(result)
}