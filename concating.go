package main 
import "fmt"


func getConcatenation(nums []int) []int {
    n := len(nums)
    ans:= make([]int, 2*n)
    for i:=0;i<n;i++{
    ans[i]=nums[i]
    ans[i+n]=nums[i]
    }
    return ans 
}
func main(){
	nums1:= []int{1,2,3,45}
	fmt.Println(getConcatenation(nums1))

	num2:= []int {3,4,5,73,11}
	fmt.Println("this is second one ",getConcatenation(num2))

}