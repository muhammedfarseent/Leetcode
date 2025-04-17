package main
import "fmt"

func averageValue(nums []int) int {
    sum,count := 0,0
    for i:=0;i<len(nums);i++{
        if nums[i]%6==0{
           sum+=nums[i]
           count++
    }
    }

    if count ==0{
        return 0
    }
    d:=sum/count
    return d
}
func main(){
	arr:= []int {1,3,6,10,12,15}
	result:=averageValue(arr)
	fmt.Println(result)
}