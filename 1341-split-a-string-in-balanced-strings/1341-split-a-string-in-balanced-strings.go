func balancedStringSplit(s string) int {
    count := 0
    left,right := 0,0
    for i:=0;i<len(s);i++{
        if s[i]=='L'{
            left ++
        }else{
            right ++
        }
        if left ==  right {
            count ++ 
            left = 0 
            right = 0
        }
    }
    return count    
}