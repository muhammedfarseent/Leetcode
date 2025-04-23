func plusOne(digits []int) []int {
    var modifier = 1 
    for i:= len(digits)-1;i>=0;i--{
        if digits[i]+ modifier>9{
            digits[i]=0
            modifier = 1
        } else {
            digits[i] += modifier
            modifier = 0
        }
    }
    if modifier == 1{
        digits = append([]int{1},digits...)
    }
    return digits
}