package main 
import ("fmt"
"strconv")


func fizzBuzz(n int) []string {
    result:= make([]string,n)


    for i:=1;i<=n;i++{
        if i%3==0 && i%5==0{
            result [i-1] = "FizzBuzz"
        }else if i%3==0{
            result [i-1]= "Fizz"
        }else if i%5==0{
        result[i-1] = "Buzz"
    } else{
        result[i-1] = strconv.Itoa(i)
    }
}
    return result 

}

func main(){
	a:=12
	fmt.Println(fizzBuzz(a))
}