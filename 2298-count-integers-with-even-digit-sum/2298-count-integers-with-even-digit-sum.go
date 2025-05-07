func countEven(num int) int {
    count := 0
    for i := 1; i <= num; i++ {
        if digitSum(i)%2 == 0 {
            count++
        }
    }
    return count
}

func digitSum(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}
