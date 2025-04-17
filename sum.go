package main 
import "fmt"


func twoSum(nums []int, target int) []int {
    n := len(nums)

    for i := 0; i < n; i++{
       for  j := i+1; j < n; j++ {
            if nums[i] + nums[j] == target {
                return []int{i,j}
            }
        }
    }
    return nil
}

func main(){
	num1:=[]int {1,4,6,8,}
    target:= 9
	result:= twoSum(num1,target)
	if result != nil{
		fmt.Println("%v",result)
	}else{
		fmt.Println("invalid ")
	}
}
