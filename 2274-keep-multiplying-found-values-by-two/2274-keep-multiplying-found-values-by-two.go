func findFinalValue(nums []int, original int) int {
    sort.Ints(nums)
    if original < nums[0]|| original >nums[len(nums)-1]{
        return original 
    }
    for i:= range nums{
        if original == nums[i]{
            original = original *2
        }
    }
    return original
    
}