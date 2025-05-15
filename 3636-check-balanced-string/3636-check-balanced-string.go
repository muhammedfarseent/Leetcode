func isBalanced(num string) bool {
    sum:= 0
    for i :=0; i<len(num);i++{
        if i% 2==0{
            sum+= int (num[i]-'0')
        }else{
            sum -= int(num[i]-'0')
        }
    }
    return sum==0
}