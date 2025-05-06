func isThree(n int) bool {
    if n == 1 || n == 2 || n == 3 {
        return false
    }

    c := 2

    for i:=2; i< n/2+1; i++{
        if n % i == 0 {
            c++
        }
    }
    return c == 3
}